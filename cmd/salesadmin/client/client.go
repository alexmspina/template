package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/alexmspina/template/api/salesadminpb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	rootCmd = &cobra.Command{
		Use:   "client",
		Short: "tests grpc file uploading",
		Long:  `Upload files to the salesadmin gRPC server.`,
	}
)

func main() {
	Execute()
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// read configuration file and initialize config variables
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/client")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	// add subcommands to the main uploader function
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(getOrdersCmd)
	rootCmd.AddCommand(getTotalRevenueCmd)
	rootCmd.AddCommand(getCustomerCountCmd)
	rootCmd.AddCommand(getMerchantCountCmd)
}

// getOrders implements the NewSalesAdminClient interface and requests
// all orders from the salesadmin api
func getOrders() {
	// creds, _ := credentials.NewClientTLSFromFile(viper.GetString("certFile"), "")

	// Set up a connection to the server.
	// conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSalesAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetAllOrders(ctx, &pb.OrdersRequest{})
	if err != nil {
		log.Fatalf("error requesting orders from server: %v", err)
	}
	log.Printf("Orders: %v", r.GetOrders())
}

// getOrdersCmd configures the getorders subcommand start
var getOrdersCmd = &cobra.Command{
	Use:   "getorders",
	Short: "request all orders from the salesadmin server",
	Long:  "Send a request to the salesadmin gRPC server for all the orders stored in the database.",
	Run: func(cmd *cobra.Command, args []string) {
		getOrders()
	},
}

// getTotalRevenue implements the NewSalesAdminClient interface and requests
// the total sales revenue from all orders stored in the database
func getTotalRevenue() {
	// creds, _ := credentials.NewClientTLSFromFile(viper.GetString("certFile"), "")

	// Set up a connection to the server.
	// conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSalesAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetTotalSalesRevenue(ctx, &pb.TotalSalesRevenueRequest{})
	if err != nil {
		log.Fatalf("could not get total sales revenue from the server: %v", err)
	}
	log.Printf("Total Sales Revenue: %v", r.GetTotalRevenue())
}

// getTotalRevenueCmd configures the getrevenue subcommand start
var getTotalRevenueCmd = &cobra.Command{
	Use:   "getrevenue",
	Short: "request total sales revenue from the salesadmin server",
	Long:  "Send a request to the salesadmin gRPC server for the total sales revenue of all the orders stored in the database.",
	Run: func(cmd *cobra.Command, args []string) {
		getTotalRevenue()
	},
}

// getCustomerCount implements the NewSalesAdminClient interface and requests
// the count of unique customers from the salesadmin api
func getCustomerCount() {
	// creds, _ := credentials.NewClientTLSFromFile(viper.GetString("certFile"), "")

	// Set up a connection to the server.
	// conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSalesAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetCustomerCount(ctx, &pb.CustomerCountRequest{})
	if err != nil {
		log.Fatalf("could not get the customer count from the server: %v", err)
	}
	log.Printf("Customer count: %v", r.GetCount())
}

// getCustomerCountCmd configures the getcustomercount subcommand start
var getCustomerCountCmd = &cobra.Command{
	Use:   "getcustomercount",
	Short: "request the count of unique customers from the salesadmin server",
	Long:  "Send a request to the salesadmin gRPC server for the count of unique customers stored in the database.",
	Run: func(cmd *cobra.Command, args []string) {
		getCustomerCount()
	},
}

// getMerchantCount implements the NewSalesAdminClient interface and requests
// the count of unique Mmerchants from the salesadmin api
func getMerchantCount() {
	// creds, _ := credentials.NewClientTLSFromFile(viper.GetString("certFile"), "")

	// Set up a connection to the server.
	// conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSalesAdminServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetMerchantCount(ctx, &pb.MerchantCountRequest{})
	if err != nil {
		log.Fatalf("could not get the merchant count from the server: %v", err)
	}
	log.Printf("Merchant count: %v", r.GetCount())
}

// getMerchantCountCmd configures the getmerchantcount subcommand start
var getMerchantCountCmd = &cobra.Command{
	Use:   "getmerchantcount",
	Short: "request the count of unique merchants from the salesadmin server",
	Long:  "Send a request to the salesadmin gRPC server for the count of unique merchants stored in the database.",
	Run: func(cmd *cobra.Command, args []string) {
		getMerchantCount()
	},
}

// upload opens the sales file csv, converts it to a slice of byte slices,
// and executes a FileUpload gRPC request
func upload() {
	// creds, _ := credentials.NewClientTLSFromFile(viper.GetString("certFile"), "")

	// Set up a connection to the server.
	// conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(viper.GetString("salesAdminAddress"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSalesAdminServiceClient(conn)

	salesFile, err := os.Open(viper.GetString("salesFile"))
	if err != nil {
		panic(err.Error())
	}

	defer salesFile.Close()

	scanner := bufio.NewScanner(salesFile)
	var salesFileBytes [][]byte

	for scanner.Scan() {
		salesFileBytes = append(salesFileBytes, []byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file := pb.File{
		FileName:  viper.GetString("salesFile"),
		FileBytes: salesFileBytes,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.FileUpload(ctx, &pb.FileUploadRequest{File: &file})
	if err != nil {
		log.Fatalf("could not upload file: %v", err)
	}
	log.Printf("Upload Success: %v", r.Result)
}

// uploadCmd configures the uploader subcommand start
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload the salesdata.csv file",
	Long:  "This uploads the salesdata.csv file to the salesadmin server",
	Run: func(cmd *cobra.Command, args []string) {
		upload()
	},
}
