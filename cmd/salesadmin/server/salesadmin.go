package main

import (
	"context"

	"github.com/alexmspina/template/internal/salesadmin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	Execute()
}

var (
	rootCmd = &cobra.Command{
		Use:   "salesadmin",
		Short: "sales data admin service",
		Long:  `Process sales data csv files.`,
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
	ctx := context.Background()
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// check for config files
	sugar.Info("checking for config files")
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/salesadmin/dev")
	viper.AddConfigPath(".")

	// set default config values
	sugar.Info("setting default config values")
	viper.SetDefault("certFile", "./server.cer.pem")
	viper.SetDefault("keyFile", "./server.key.pem")
	viper.SetDefault("caFile", "./salesadmin-ca-1.crt")
	viper.SetDefault("port", ":9090")
	viper.SetDefault("postgresCertFile", "./postgres.cer.pem")
	viper.SetDefault("postgresKeyFile", "./postgres.key.pem")
	viper.SetDefault("postgresHost", "postgres")
	viper.SetDefault("postgresPort", 5432)
	viper.SetDefault("postgresUser", "salesadmin")
	viper.SetDefault("postgresDB", "salesadmin")
	viper.SetDefault("postgresPassword", "salesadmin")

	err := viper.ReadInConfig()
	if err != nil {
		sugar.Debug(" error reading config values")
	} else {
		for k, v := range viper.AllSettings() {
			sugar.Infof("Config parameter %v: %v", k, v)
		}
	}

	// initialize database
	sugar.Info("creating orders table in the database")
	err = salesadmin.CreateOrdersTable(ctx)
	if err != nil {
		sugar.Debugf("error %v creating orders table in the database", err)
	}

	sugar.Info("configuring subcommands")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the salesadmin service",
	Long:  "This command starts the salesadmin gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		salesadmin.Start()
	},
}
