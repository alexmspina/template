package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	Execute()
}

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "dse",
		Short: "Dynamic sensor enrollment",
		Long:  `Dynamically link new sensors to the network.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello, dse!",
	Long:  "This command is a viper config test.",
	Run: func(cmd *cobra.Command, args []string) {
		hello()
	},
}

func hello() {
	fmt.Println(viper.GetString("hello"))
}

func init() {
	viper.SetDefault("hello", "Hello, dse!")

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/dse")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(konnectCmd)
}

func konnect() {
	producerCert, err := tls.LoadX509KeyPair(viper.GetString("producerCertFilePath"), viper.GetString("producerKeyFilePath"))
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
	// Load CA cert
	caCert, err := ioutil.ReadFile(viper.GetString("portappCAFilePath"))
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := tls.Config{
		Certificates: []tls.Certificate{producerCert},
		RootCAs:      caCertPool,
	}

	dialer := &kafka.Dialer{
		ClientID: "connection test",
		TLS:      &tlsConfig,
	}

	// dialer := &kafka.Dialer{
	// 	ClientID: "connection test",
	// }

	conn, err := dialer.Dial("tcp", "broker1:19092")
	if err != nil {
		fmt.Println(fmt.Errorf("error connecting to kafka %v", err))
	}
	fmt.Println(conn)

	conn.Close()
}

var konnectCmd = &cobra.Command{
	Use:   "konnect",
	Short: "Connect to Kafka from dse service",
	Long:  "This command tests dse's ability to connect to Kafka",
	Run: func(cmd *cobra.Command, args []string) {
		konnect()
	},
}
