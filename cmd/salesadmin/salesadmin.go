package main

import (
	"fmt"
	"github.com/alexmspina/template/internal/salesadmin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	return rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/salesadmin")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file %s", err))
	}

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
