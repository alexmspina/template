package main

import (
	"github.com/alexmspina/template/internal/musicworld"
	"github.com/spf13/viper"
)

func main() {
	musicworld.Execute()
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
}
