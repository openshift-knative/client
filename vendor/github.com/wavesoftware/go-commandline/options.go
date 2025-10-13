package commandline

import (
	"io"

	"github.com/spf13/cobra"
)

// Option is used to configure an App.
type Option func(*App)

// WithArgs creates an option which sets args.
// Deprecated: use WithCommand instead.
func WithArgs(args ...string) Option {
	return WithCommand(func(cmd *cobra.Command) {
		cmd.SetArgs(args)
	})
}

// WithInput creates an option witch sets os.Stdin.
// Deprecated: use WithCommand instead.
func WithInput(in io.Reader) Option {
	return WithCommand(func(cmd *cobra.Command) {
		cmd.SetIn(in)
	})
}

// WithOutput creates an option witch sets os.Stdout and os.Stderr.
// Deprecated: use WithCommand instead.
func WithOutput(out io.Writer) Option {
	return WithCommand(func(cmd *cobra.Command) {
		cmd.SetOut(out)
		cmd.SetErr(out)
	})
}

// WithCommand will allow one to change the cobra.Command.
func WithCommand(fn func(cmd *cobra.Command)) Option {
	return func(app *App) {
		fn(app.root)
	}
}

// WithExit creates an option which sets the exit function.
func WithExit(fn func(code int)) Option {
	return func(app *App) {
		app.Exit = fn
	}
}
