package musicworld

import "github.com/spf13/cobra"

func init() {
	// add subcommands to root command, musicworld
	rootCmd.AddCommand(SpinCmd)
}

var (
	rootCmd = &cobra.Command{
		Use:   "musicworld COMMAND",
		Short: "Mix it up!",
		Long: `
A toolset for building interactive
audio-visual experiences with
WASM.

Commands:
	spin		start a service`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
