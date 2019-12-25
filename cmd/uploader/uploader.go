package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	pb "github.com/alexmspina/template/api/salesadminpb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "uploader",
		Short: "tests grpc file uploading",
		Long:  `Upload files to the salesprocessor service.`,
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
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/uploader")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	rootCmd.AddCommand(startCmd)
}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

func start() {
	creds, _ := credentials.NewClientTLSFromFile(viper.GetString("certFile"), "")
	fmt.Printf("salesProcessor address %v\n\n", viper.GetString("salesProcessorAddress"))
	// Set up a connection to the server.
	conn, err := grpc.Dial(viper.GetString("salesProcessorAddress"), grpc.WithTransportCredentials(creds))
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

	reader := bufio.NewReader(salesFile)

	var chunk []byte
	var eol bool
	var salesFileBytes [][]byte

	for {
		if chunk, eol, err = reader.ReadLine(); err != nil {
			break
		}

		if !eol {
			salesFileBytes = append(salesFileBytes, chunk)
		}
	}

	if err == io.EOF {
		err = nil
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

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the uploader client",
	Long:  "This command starts the file uploader gRPC client",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}
