package salesadmin

import (
	"context"
	"log"
	"net"

	pb "github.com/alexmspina/template/api/salesadminpb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSalesAdminServiceServer
}

// Start the SalesAdminService gRPC server
func Start() {
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

// FileUpload implements the FileUpload methods of the SalesAdminServiceServer interface
func (s *server) FileUpload(ctx context.Context, in *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	file := in.GetFile()
	fileBytes := file.FileBytes
	orders, err := ParseSalesFile(ctx, fileBytes)
	if err != nil {
		return nil, err
	}

	err = InsertOrders(ctx, orders)
	if err != nil {
		return nil, err
	}

	return &pb.FileUploadResponse{Result: true}, nil
}

// ServeOrders implements the Orders methods of the SalesAdminServiceServer
func (s *server) GetAllOrders(ctx context.Context, in *pb.OrdersRequest) (*pb.OrdersResponse, error) {
	orders, err := RunQueryAllOrders(ctx, queryAllOrders)
	if err != nil {
		return nil, err
	}

	var allOrders []*pb.Order
	for _, order := range orders {
		pbOrder := pb.Order{
			OrderId:         order.OrderID,
			CustomerName:    order.CustomerName,
			ItemPrice:       order.ItemPrice,
			ItemDescription: order.ItemDescription,
			Quantity:        order.Quantity,
			MerchantName:    order.MerchantName,
			MerchantAddress: order.MerchantAddress,
		}
		allOrders = append(allOrders, &pbOrder)
	}

	response := pb.OrdersResponse{
		Orders: allOrders,
	}

	return &response, nil
}

// ServeTotalRevenue implements the Orders methods of the SalesAdminServiceServer
func (s *server) GetTotalSalesRevenue(ctx context.Context, in *pb.TotalSalesRevenueRequest) (*pb.TotalSalesRevenueResponse, error) {
	orders, err := RunQueryTotalRevenue(ctx, queryRevenueTotal)
	if err != nil {
		return nil, err
	}

	var totalRevenue float64
	for _, order := range orders {
		price := order.ItemPrice
		quantity := float64(order.Quantity)
		totalRevenue = float64(totalRevenue) + price*quantity
	}

	response := pb.TotalSalesRevenueResponse{
		TotalRevenue: totalRevenue,
	}

	return &response, nil
}
