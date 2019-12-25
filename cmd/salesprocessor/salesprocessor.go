package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/alexmspina/template/api/salesadminpb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	Execute()
}

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "salesprocessor",
		Short: "sales data processing service",
		Long:  `Process sales data csv files.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/salesprocessor")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file %s", err))
	}

	rootCmd.AddCommand(startCmd)
}

type server struct {
	pb.UnimplementedSalesAdminServiceServer
}

// SayHello implements helloworld.GreeterServer
// func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
// 	log.Printf("Received: %v", in.GetName())
// 	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
// }

func (s *server) FileUpload(ctx context.Context, in *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	file := in.GetFile()
	log.Printf("Received file: %v", file.FileName)
	log.Printf("File bytes: \n%v", file.FileBytes)
	return &pb.FileUploadResponse{Result: true}, nil
}

func start() {
	creds, _ := credentials.NewServerTLSFromFile(viper.GetString("certFile"), viper.GetString("keyFile"))
	lis, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSalesAdminServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the sales processor service",
	Long:  "This command starts the sales processing gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}
