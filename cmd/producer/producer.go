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
		Use:   "producer",
		Short: "Producer test program",
		Long:  `A simple program to produce messages from Kafka.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello, producer!",
	Long:  "This command is a viper config test.",
	Run: func(cmd *cobra.Command, args []string) {
		hello()
	},
}

func hello() {
	fmt.Println(viper.GetString("hello"))
}

var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce Kafka messages",
	Long:  "This program tests producing messages on a Kafka topic",
	Run: func(cmd *cobra.Command, args []string) {
		produce()
	},
}

func produce() {
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
		ClientID: "produce messages over ssl",
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

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)

	conn.Close()
}

func init() {
	viper.SetDefault("hello", "Hello, producer!")

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/producer")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(produceCmd)
}
