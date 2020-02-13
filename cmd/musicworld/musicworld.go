package main

import (
	"github.com/alexmspina/template/internal/musicworld"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	Execute()
}

var (
	rootCmd = &cobra.Command{
		Use:   "musicworld COMMAND",
		Short: "Mix it up!",
		Long: `A toolset for building interactive
audio-visual experiences with
WASM.

Commands:
  spin		start a service`,
	}
)

// Execute executes the root command.
func Execute() error {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("executing root command")
	return rootCmd.Execute()
}

func init() {
	// configure init context
	// ctx := context.Background()
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// check for config files
	sugar.Info("checking for config files")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/musicworld/dev")
	viper.AddConfigPath(".")

	// set default config values
	sugar.Info("setting default config values")
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

	err := viper.ReadInConfig()
	if err != nil {
		sugar.Debug(" error reading config values")
	} else {
		for k, v := range viper.AllSettings() {
			sugar.Infof("Config parameter %v: %v", k, v)
		}
	}

	// initialize database
	// sugar.Info("this is where the postgres init script will go")
	// err = salesadmin.CreateOrdersTable(ctx)
	// if err != nil {
	//     	sugar.Debugf("error %v creating orders table in the database", err)
	// }

	sugar.Info("configuring subcommands")
	rootCmd.AddCommand(spinCmd)
}

var spinCmd = &cobra.Command{
	Use:   "spin SERIVCE",
	Short: "spin up some vinyl!",
	Long: `Start a service

Services:
  devcontainer		a general rust-alpine image with terminal based dev tools`,
	Run: func(cmd *cobra.Command, args []string) {
		musicworld.Spin()
	},
}
