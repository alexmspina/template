package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

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
		Use:   "pgclient",
		Short: "a golang postgres client service",
		Long:  `Run a golang postgres client service.`,
	}
)

const (
	host     = "postgres"
	port     = 5432
	user     = "pad"
	dbname   = "paddb"
	password = "padsword"
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Hello, pgclient!",
	Long:  "This command is a viper config test.",
	Run: func(cmd *cobra.Command, args []string) {
		hello()
	},
}

func hello() {
	fmt.Println(viper.GetString("hello"))
}

func init() {
	viper.SetDefault("hello", "Hello, pgclient!")

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/pgclient")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(connectCmd)
}

func connect() {
	// producerCert, err := tls.LoadX509KeyPair(viper.GetString("pgclientCertFilePath"), viper.GetString("pgclientKeyFilePath"))
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("%v", err))
	// }
	// // Load CA cert
	// caCert, err := ioutil.ReadFile(viper.GetString("portappCAFilePath"))
	// if err != nil {
	// 	fmt.Println(fmt.Errorf("%v", err))
	// }
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// tlsConfig := tls.Config{
	// 	Certificates: []tls.Certificate{producerCert},
	// 	RootCAs:      caCertPool,
	// }
	fmt.Println("hello")
	sslcert := viper.GetString("postgresCertFilePath")
	sslkey := viper.GetString("postgresKeyFilePath")
	sslrootcert := viper.GetString("rootCaFilePath")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslkey=%s sslcert=%s sslrootcert=%s sslmode=verify-ca", host, port, user, password, dbname, sslkey, sslcert, sslrootcert)
	// psqlInfo := "user=pad dbname=paddb port=5432 host=postgres password=padsword sslmode=disable"

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("This is the first panic")
		panic(err)
	}
	defer db.Close()

	fmt.Println("success", db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a Postgres server from the pgclient service",
	Long:  "This command tests pgclient's ability to connect to Postgres",
	Run: func(cmd *cobra.Command, args []string) {
		connect()
	},
}
