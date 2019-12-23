package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"time"

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
		Use:   "consumer",
		Short: "Consumer test program",
		Long:  `A simple program to consumer messages from Kafka.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello, consumer!",
	Long:  "This command is a viper config test.",
	Run: func(cmd *cobra.Command, args []string) {
		hello()
	},
}

func hello() {
	fmt.Println(viper.GetString("hello"))
}

var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume messages on a Kafka topic.",
	Long:  "This program tests the ability of the dev container to consume messages from Kafka.",
	Run: func(cmd *cobra.Command, args []string) {
		consume()
	},
}

func consume() {
	consumerCert, err := tls.LoadX509KeyPair(viper.GetString("consumerCertFilePath"), viper.GetString("consumerKeyFilePath"))
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
		Certificates: []tls.Certificate{consumerCert},
		RootCAs:      caCertPool,
	}

	dialer := &kafka.Dialer{
		ClientID: "consume messages over ssl",
		TLS:      &tlsConfig,
	}

	// dialer := &kafka.Dialer{
	// 	ClientID: "connection test",
	// }

	// to produce messages
	topic := "my-topic"
	partition := 0

	// to connect to leader port needs to be 19092
	conn, err := dialer.DialLeader(context.Background(), "tcp", "broker1:19092", topic, partition)
	if err != nil {
		fmt.Println(fmt.Errorf("error connecting to kafka leader %v", err))
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}

	batch.Close()
	conn.Close()
}

func init() {
	viper.SetDefault("hello", "Hello, consumer!")

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/consumer")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(consumeCmd)
}
