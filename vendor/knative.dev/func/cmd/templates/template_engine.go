// Copyright © 2020 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"golang.org/x/term"
)

type templateEngine struct {
	RootCmd *cobra.Command
	CommandGroups
	functions template.FuncMap
}

// Create new template engine
func newTemplateEngine(rootCmd *cobra.Command, g CommandGroups, extraFunctions *template.FuncMap) templateEngine {
	engine := templateEngine{
		RootCmd:       rootCmd,
		CommandGroups: g,
	}
	engine.functions = engine.templateFunctions()
	if extraFunctions != nil {
		for name, function := range *extraFunctions {
			engine.functions[name] = function
		}
	}
	return engine
}

func (e templateEngine) usageFunc() func(*cobra.Command) error {
	return func(c *cobra.Command) error {
		return e.fillTemplate("usage", c, usageTemplate())
	}
}

func (e templateEngine) helpFunc() func(*cobra.Command, []string) {
	return func(c *cobra.Command, s []string) {
		err := e.fillTemplate("help", c, helpTemplate(c.Long))
		if err != nil {
			c.Println(err)
		}
	}
}

func (e templateEngine) fillTemplate(name string, c *cobra.Command, templ string) error {
	t := template.New(name)
	t.Funcs(e.functions)
	_, err := t.Parse(templ)
	if err != nil {
		fmt.Fprintf(c.ErrOrStderr(), "\nINTERNAL: >>>>> %v\n", err)
		return err
	}
	return t.Execute(c.OutOrStdout(), c)
}

// ======================================================================================
// Template helper functions

func (e templateEngine) templateFunctions() template.FuncMap {
	return template.FuncMap{
		"cmdGroupsString":   e.cmdGroupsString,
		"subCommandsString": e.subCommandsString,
		"useLine":           useLine,
		"visibleFlags":      visibleFlags,
		"rpad":              rpad,
		"rootCmdName":       e.rootCmdName,
		"rootCmdUse":        e.rootCmdUse,
		"isRootCmd":         e.isRootCmd,
		"flagsUsages":       flagsUsagesCobra, // or use flagsUsagesKubectl for kubectl like flag styles
		"trim":              strings.TrimSpace,
		"trimRight":         func(s string) string { return strings.TrimRightFunc(s, unicode.IsSpace) },
		"trimLeft":          func(s string) string { return strings.TrimLeftFunc(s, unicode.IsSpace) },
		"execTemplate":      e.executeTemplate,
	}
}

func (e templateEngine) executeTemplate(tbody string, data any) (string, error) {
	t, err := template.New("").Funcs(e.templateFunctions()).Parse(tbody)
	if err != nil {
		return "", err
	}
	buf := &strings.Builder{}
	err = t.Execute(buf, data)
	return buf.String(), err
}

func (e templateEngine) cmdGroupsString() string {
	groups := make([]string, 0, len(e.CommandGroups))
	for _, cmdGroup := range e.CommandGroups {
		groups = append(groups, formatCommandGroup(cmdGroup))
	}
	return strings.Join(groups, "\n\n")
}

func (e templateEngine) subCommandsString(c *cobra.Command) string {
	return formatCommandGroup(CommandGroup{
		Header:   "Available Commands:",
		Commands: c.Commands(),
	})
}

func (e templateEngine) rootCmdName() string {
	return e.RootCmd.CommandPath()
}

func (e templateEngine) rootCmdUse() string {
	return e.RootCmd.Use
}

func (e templateEngine) isRootCmd(c *cobra.Command) bool {
	return e.RootCmd == c
}

func visibleFlags(c *cobra.Command) *flag.FlagSet {
	ret := flag.NewFlagSet("filtered", flag.ContinueOnError)
	local := c.LocalFlags()
	persistent := c.PersistentFlags()
	local.VisitAll(func(flag *flag.Flag) {
		if flag.Name != "help" && persistent.Lookup(flag.Name) == nil {
			ret.AddFlag(flag)
		}
	})
	return ret
}

func useLine(c *cobra.Command) string {
	var useLine string
	var suffix string
	if c.HasParent() {
		useLine = c.Parent().CommandPath() + " " + c.Use
		suffix = "[flags]"
	} else {
		useLine = c.Use
		suffix = "[command]"
	}
	if c.HasFlags() && !strings.Contains(useLine, suffix) {
		useLine += " " + suffix
	}
	return useLine
}

func formatCommandGroup(cmdGroup CommandGroup) string {
	cmds := []string{cmdGroup.Header}
	for _, cmd := range cmdGroup.Commands {
		if cmd.IsAvailableCommand() {
			cmds = append(cmds, "  "+rpad(cmd.Name(), cmd.NamePadding())+" "+cmd.Short)
		}
	}
	return strings.Join(cmds, "\n")
}

func rpad(s string, padding int) string {
	t := fmt.Sprintf("%%-%ds", padding)
	return fmt.Sprintf(t, s)
}

// flagsUsagesCobra formats flags in Cobra style
func flagsUsagesCobra(f *flag.FlagSet) string {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err == nil {
		return f.FlagUsagesWrapped(width)
	} else {
		return f.FlagUsages()
	}
}

// flagsUsagesKubectl formats the flags like kubectl does
func flagsUsagesKubectl(f *flag.FlagSet) string {
	x := new(bytes.Buffer)

	f.VisitAll(func(flag *flag.Flag) {
		if flag.Hidden {
			return
		}
		format := "--%s=%s: %s\n"

		if flag.Value.Type() == "string" {
			format = "--%s='%s': %s\n"
		}

		if len(flag.Shorthand) > 0 {
			format = "  -%s, " + format
		} else {
			format = "   %s   " + format
		}

		fmt.Fprintf(x, format, flag.Shorthand, flag.Name, flag.DefValue, flag.Usage)
	})

	return x.String()
}
