package salesadmin

import (
	"context"
	"net"

	pb "github.com/alexmspina/template/api/salesadminpb"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedSalesAdminServiceServer
}

// Start the SalesAdminService gRPC server
func Start() {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("configuring gRPC credentials")
	creds, _ := credentials.NewServerTLSFromFile(viper.GetString("certFile"), viper.GetString("keyFile"))

	sugar.Infof("configuring gRPC listening on port %v", viper.GetString("port"))
	lis, err := net.Listen("tcp", viper.GetString("port"))
	if err != nil {
		sugar.Fatalf("failed to listen: %v", err)
	}

	sugar.Info("starting SalesAdminServer")
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterSalesAdminServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		sugar.Fatalf("failed to serve: %v", err)
	}
}

// FileUpload implements the FileUpload methods of the SalesAdminServiceServer interface
func (s *server) FileUpload(ctx context.Context, in *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("reading FileUploadRequest FileBytes payload")
	file := in.GetFile()
	fileBytes := file.FileBytes
	orders, err := ParseSalesFile(ctx, fileBytes)
	if err != nil {
		return nil, err
	}

	sugar.Info("calling InsertOrders with FileBytes payload")
	err = InsertOrders(ctx, orders)
	if err != nil {
		return nil, err
	}

	return &pb.FileUploadResponse{Result: true}, nil
}

// ServeOrders implements the Orders methods of the SalesAdminServiceServer
func (s *server) GetAllOrders(ctx context.Context, in *pb.OrdersRequest) (*pb.OrdersResponse, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("calling RunQueryAllOrders method")
	orders, err := RunQueryAllOrders(ctx, queryAllOrders)
	if err != nil {
		return nil, err
	}

	sugar.Info("unmarshalling orders into Order structs")
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

	sugar.Info("creating OrderResponse payload for gRPC GetAllOrders")
	response := pb.OrdersResponse{
		Orders: allOrders,
	}

	return &response, nil
}

// ServeTotalRevenue implements the Orders methods of the SalesAdminServiceServer
func (s *server) GetTotalSalesRevenue(ctx context.Context, in *pb.TotalSalesRevenueRequest) (*pb.TotalSalesRevenueResponse, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("calliong RunQueryTotalRevenue method")
	orders, err := RunQueryTotalRevenue(ctx, queryRevenueTotal)
	if err != nil {
		return nil, err
	}

	sugar.Info("calculating total revenue")
	var totalRevenue float64
	for _, order := range orders {
		price := order.ItemPrice
		quantity := float64(order.Quantity)
		totalRevenue = float64(totalRevenue) + price*quantity
	}

	sugar.Info("creating TotalSalesRevenueResponse for gRPC method GetTotalSalesRevenue")
	response := pb.TotalSalesRevenueResponse{
		TotalRevenue: totalRevenue,
	}

	return &response, nil
}
