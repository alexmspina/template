package main

import (
	"fmt"

	"github.com/alexmspina/template/internal/musicworld"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	Execute()
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

func init() {
	// search for config files
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/musicworld/dev")
	viper.AddConfigPath(".")

	// set default config values
	viper.SetDefault("certFile", "./server.cer.pem")
	viper.SetDefault("keyFile", "./server.key.pem")
	viper.SetDefault("caFile", "./musicworld-ca-1.crt")
	viper.SetDefault("port", ":9090")
	viper.SetDefault("postgresCertFile", "./postgres.cer.pem")
	viper.SetDefault("postgresKeyFile", "./postgres.key.pem")
	viper.SetDefault("postgresHost", "postgres")
	viper.SetDefault("postgresPort", 5432)
	viper.SetDefault("postgresUser", "musicworld")
	viper.SetDefault("postgresDB", "musicworld")
	viper.SetDefault("postgresPassword", "musicworld")

	// add subcommands to root command, musicworld
	rootCmd.AddCommand(spinCmd)
}

var spinCmd = &cobra.Command{
	Use:   "spin SERIVCE",
	Short: "spin up some vinyl!",
	Long: `
Start a service

Services:
  devcontainer		a general rust-alpine image with terminal based dev tools`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(`
"musicworld spin" requires exactly one argument.

See 'musicworld spin --help'.

Usage: musicworld spin SERVICE
`)
		} else {
			musicworld.Spin(args[0])
		}
	},
}
