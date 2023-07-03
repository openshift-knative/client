package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ory/viper"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/term"
	"k8s.io/apimachinery/pkg/util/sets"
	"knative.dev/client-pkg/pkg/util"

	"knative.dev/func/cmd/templates"
	"knative.dev/func/pkg/config"
	fn "knative.dev/func/pkg/functions"
	"knative.dev/func/pkg/k8s"
)

type RootCommandConfig struct {
	Name string // usually `func` or `kn func`
	Version
	NewClient ClientFactory
}

// NewRootCmd creates the root of the command tree defines the command name, description, globally
// available flags, etc.  It has no action of its own, such that running the
// resultant binary with no arguments prints the help/usage text.
func NewRootCmd(cfg RootCommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   cfg.Name,
		Short: fmt.Sprintf("%s manages Knative Functions", cfg.Name),
		Long: fmt.Sprintf(`%s is the command line interface for managing Knative Function resources

	Create a new Node.js function in the current directory:
	{{.Use}} create --language node myfunction

	Deploy the function using Docker hub to host the image:
	{{.Use}} deploy --registry docker.io/alice

Learn more about Functions:  https://knative.dev/docs/functions/
Learn more about Knative at: https://knative.dev`, cfg.Name),

		DisableAutoGenTag: true, // no docs header
		SilenceUsage:      true, // no usage dump on error
		SilenceErrors:     true, // we explicitly handle errors in Execute()
	}

	// Environment Variables
	// Evaluated first after static defaults, set all flags to be associated with
	// a version prefixed by "FUNC_"
	viper.AutomaticEnv()       // read in environment variables for FUNC_<flag>
	viper.SetEnvPrefix("func") // ensure that all have the prefix

	// Client
	// Use the provided ClientFactory or default to NewClient
	newClient := cfg.NewClient
	if newClient == nil {
		newClient = NewClient
	}

	// Grouped commands
	groups := templates.CommandGroups{
		{
			Header: "Primary Commands:",
			Commands: []*cobra.Command{
				NewCreateCmd(newClient),
				NewDescribeCmd(newClient),
				NewDeployCmd(newClient),
				NewDeleteCmd(newClient),
				NewListCmd(newClient),
			},
		},
		{
			Header: "Development Commands:",
			Commands: []*cobra.Command{
				NewRunCmd(newClient),
				NewInvokeCmd(newClient),
				NewBuildCmd(newClient),
			},
		},
		{
			Header: "System Commands:",
			Commands: []*cobra.Command{
				NewConfigCmd(defaultLoaderSaver, newClient),
				NewLanguagesCmd(newClient),
				NewTemplatesCmd(newClient),
				NewRepositoryCmd(newClient),
			},
		},
		{
			Header: "Other Commands:",
			Commands: []*cobra.Command{
				NewCompletionCmd(),
				NewVersionCmd(cfg.Version),
			},
		},
	}

	// Add all commands to the root command, and initialize
	groups.AddTo(cmd)
	groups.SetRootUsage(cmd, nil)

	return cmd
}

// Helpers
// ------------------------------------------

// registry to use is that provided as --registry or FUNC_REGISTRY.
// If not provided, global configuration determines the default to use.
func registry() string {
	if r := viper.GetString("registry"); r != "" {
		return r
	}
	cfg, _ := config.NewDefault()
	return cfg.RegistryDefault()
}

// effectivePath to use is that which was provided by --path or FUNC_PATH.
// Manually parses flags such that this can be used during (cobra/viper) flag
// definition (prior to parsing).
func effectivePath() (path string) {
	var (
		env = os.Getenv("FUNC_PATH")
		fs  = pflag.NewFlagSet("", pflag.ContinueOnError)
		p   = fs.StringP("path", "p", "", "")
	)
	fs.SetOutput(io.Discard)
	fs.ParseErrorsWhitelist.UnknownFlags = true // wokeignore:rule=whitelist
	// Preparsing flags intentionally ignores errors because this is intended
	// to be an opportunistic parse of the path flags, with actual validation of
	// flags taking place later in the instantiation process by the cobra pkg.
	_ = fs.Parse(os.Args[1:])
	if env != "" {
		path = env
	}
	if *p != "" {
		path = *p
	}
	return path
}

// interactiveTerminal returns whether or not the currently attached process
// terminal is interactive.  Used for determining whether or not to
// interactively prompt the user to confirm default choices, etc.
func interactiveTerminal() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

// bindFunc which conforms to the cobra PreRunE method signature
type bindFunc func(*cobra.Command, []string) error

// bindEnv returns a bindFunc that binds env vars to the named flags.
func bindEnv(flags ...string) bindFunc {
	return func(cmd *cobra.Command, args []string) (err error) {
		for _, flag := range flags {
			if err = viper.BindPFlag(flag, cmd.Flags().Lookup(flag)); err != nil {
				return
			}
		}
		viper.AutomaticEnv()       // read in environment variables for FUNC_<flag>
		viper.SetEnvPrefix("func") // ensure that all have the prefix
		return
	}
}

// deriveName returns the explicit value (if provided) or attempts to derive
// from the given path.  Path is defaulted to current working directory, where
// a function configuration, if it exists and contains a name, is used.
func deriveName(explicitName string, path string) string {
	// If the name was explicitly provided, use it.
	if explicitName != "" {
		return explicitName
	}

	// If the directory at path contains an initialized function, use the name therein
	f, err := fn.NewFunction(path)
	if err == nil && f.Name != "" {
		return f.Name
	}

	return ""
}

// deriveNameAndAbsolutePathFromPath returns resolved function name and absolute path
// to the function project root. The input parameter path could be one of:
// 'relative/path/to/foo', '/absolute/path/to/foo', 'foo' or ”.
func deriveNameAndAbsolutePathFromPath(path string) (string, string) {
	var absPath string

	// If path is not specified, we would like to use current working dir
	if path == "" {
		path = cwd()
	}

	// Expand the passed function name to its absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", ""
	}

	// Get the name of the function, which equals to name of the current directory
	pathParts := strings.Split(strings.TrimRight(path, string(os.PathSeparator)), string(os.PathSeparator))
	return pathParts[len(pathParts)-1], absPath
}

// deriveImage returns the same image name which will be used.
// I.e. if the explicit name is empty, derive one from the configured registry
// (registry plus username) and the function's name.
//
// This is calculated preemptively here in the CLI (prior to invoking the
// client), only in order to provide information to the user via the prompt.
// The client will calculate this same value if the image override is not
// provided.
//
// Derivation logic:
// deriveImage attempts to arrive at a final, full image name:
//
//	format:  [registry]/[username]/[functionName]:[tag]
//	example: quay.io/myname/my.function.name:tag.
//
// Registry can optionally be omitted, in which case DefaultRegistry
// will be prepended.
//
// If the image flag is provided, this value is used directly (the user supplied
// --image or $FUNC_IMAGE).  Otherwise, the function at 'path' is loaded, and
// the Image name therein is used (i.e. it was previously calculated).
// Finally, the default registry is used, which is prepended to the function
// name, and appended with ':latest':
func deriveImage(explicitImage, defaultRegistry, path string) string {
	if explicitImage != "" {
		return explicitImage // use the explicit value provided.
	}
	f, err := fn.NewFunction(path)
	if err != nil {
		return "" // unable to derive due to load error (uninitialized?)
	}
	if f.Image != "" {
		return f.Image // use value previously provided or derived.
	}
	// Use the func system's derivation logic.
	// Errors deriving result in an empty return
	derivedValue, _ := f.ImageName()
	return derivedValue
}

func envFromCmd(cmd *cobra.Command) (*util.OrderedMap, []string, error) {
	if cmd.Flags().Changed("env") {
		env, err := cmd.Flags().GetStringArray("env")
		if err != nil {
			return nil, []string{}, fmt.Errorf("Invalid --env: %w", err)
		}
		return util.OrderedMapAndRemovalListFromArray(env, "=")
	}
	return util.NewOrderedMap(), []string{}, nil
}

func mergeEnvs(envs []fn.Env, envToUpdate *util.OrderedMap, envToRemove []string) ([]fn.Env, int, error) {
	updated := sets.NewString()

	var counter int

	for i := range envs {
		if envs[i].Name != nil {
			value, present := envToUpdate.GetString(*envs[i].Name)
			if present {
				envs[i].Value = &value
				updated.Insert(*envs[i].Name)
				counter++
			}
		}
	}

	it := envToUpdate.Iterator()
	for name, value, ok := it.NextString(); ok; name, value, ok = it.NextString() {
		if !updated.Has(name) {
			n := name
			v := value
			envs = append(envs, fn.Env{Name: &n, Value: &v})
			counter++
		}
	}

	for _, name := range envToRemove {
		for i, envVar := range envs {
			if *envVar.Name == name {
				envs = append(envs[:i], envs[i+1:]...)
				counter++
				break
			}
		}
	}

	errMsg := fn.ValidateEnvs(envs)
	if len(errMsg) > 0 {
		return []fn.Env{}, 0, fmt.Errorf(strings.Join(errMsg, "\n"))
	}

	return envs, counter, nil
}

// addConfirmFlag ensures common text/wording when the --path flag is used
func addConfirmFlag(cmd *cobra.Command, dflt bool) {
	cmd.Flags().BoolP("confirm", "c", dflt, "Prompt to confirm options interactively ($FUNC_CONFIRM)")
}

// addPathFlag ensures common text/wording when the --path flag is used
func addPathFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("path", "p", "", "Path to the function.  Default is current directory ($FUNC_PATH)")
}

// addVerboseFlag ensures common text/wording when the --path flag is used
func addVerboseFlag(cmd *cobra.Command, dflt bool) {
	cmd.Flags().BoolP("verbose", "v", false, "Print verbose logs ($FUNC_VERBOSE)")
}

// cwd returns the current working directory or exits 1 printing the error.
func cwd() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Unable to determine current working directory: %v", err))
	}
	return cwd
}

type Version struct {
	// Date of compilation
	Date string
	// Version tag of the git commit, or 'tip' if no tag.
	Vers string
	// Hash of the currently active git commit on build.
	Hash string
	// Verbose printing enabled for the string representation.
	Verbose bool
}

// Return the stringification of the Version struct, which takes into account
// the verbosity setting.
func (v Version) String() string {
	if v.Verbose {
		return v.StringVerbose()
	}

	// Ensure that the value returned is parseable as a semver, with the special
	// value v0.0.0 as the default indicating there is no version information
	// available.
	if strings.HasPrefix(v.Vers, "v") {
		// TODO: this is the naive approach, perhaps consider actually parse it
		// using the semver lib
		return v.Vers
	}

	// Any non-semver value is invalid, and thus indistinguishable from a
	// nonexistent version value, so the default zero value of v0.0.0 is used.
	return "v0.0.0"
}

// StringVerbose returns the verbose version of the version stringification.
// The format returned is [semver]-[hash]-[date] where the special value
// 'v0.0.0' and 'source' are used when version is not available and/or the
// libray has been built from source, respectively.
func (v Version) StringVerbose() string {
	var (
		vers = v.Vers
		hash = v.Hash
		date = v.Date
	)
	if vers == "" {
		vers = "v0.0.0"
	}
	if hash == "" {
		hash = "source"
	}
	if date == "" {
		date = time.Now().Format(time.RFC3339)
	}
	funcVersion := fmt.Sprintf("%s-%s-%s", vers, hash, date)
	return fmt.Sprintf("Version: %s\n"+
		"SocatImage: %s\n"+
		"TarImage: %s", funcVersion,
		k8s.SocatImage,
		k8s.TarImage)
}

// surveySelectDefault returns 'value' if defined and exists in 'options'.
// Otherwise, options[0] is returned if it exists.  Empty string otherwise.
//
// Usage Example:
//
//	languages := []string{ "go", "node", "rust" },
//	survey.Select{
//	  Options: options,
//	  Default: surveySelectDefaut(cfg.Language, languages),
//	}
//
// Summary:
//
// This protects against an incorrectly initialized survey.Select when the user
// has provided a nonexistant option (validation is handled elsewhere) or
// when a value is required but there exists no defaults (no default value on
// the associated flag).
//
// Explanation:
//
// The above example chooses the default for the Survey (--confirm) question
// in a way that works with user-provided flag and environment variable values.
//
//	`cfg.Language` is the current value set in the config struct, which is
//	   populated from (in ascending order of precedence):
//	   static flag default, associated environment variable, or command flag.
//	`languages` are the options which are being used by the survey select.
//
// This cascade allows for the Survey questions to be properly pre-initialzed
// with their associated environment variables or flags.  For example,
// A user whose default language is set to 'node' using the global environment
// variable FUNC_LANGUAGE will have that option pre-selected when running
// `func create -c`.
//
// The 'survey' package expects the value of the Default member to exist
// in the 'Options' member.  This is not possible when user-provided data is
// allowed for the default, hence this logic is necessary.
//
// For example, when the user is using prompts (--confirm) to select from a set
// of options, but the associated flag either has an unrecognized value, or no
// value at all, without this logic the resulting select prompt would be
// initialized with this as the default value, and the act of what appears to
// be choose the first option displayed does not overwrite the invalid default.
// It could perhaps be argued this is a shortcoming in the survey package, but
// it is also clearly an error to provide invalid data for a default.
func surveySelectDefault(value string, options []string) string {
	for _, v := range options {
		if value == v {
			return v // The provided value is acceptable
		}
	}
	if len(options) > 0 {
		return options[0] // Sync with the option which will be shown by the UX
	}
	// Either the value is not an option or there are no options.  Either of
	// which should fail proper validation
	return ""
}
