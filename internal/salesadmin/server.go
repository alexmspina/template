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

// GetAllOrders implements the Orders methods of the SalesAdminServiceServer
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

// GetCustomerCount implements the Orders method of the SalesAdminServiceServer
func (s *server) GetCustomerCount(ctx context.Context, in *pb.CustomerCountRequest) (*pb.CustomerCountResponse, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("calling RunQueryCustomerNames method")
	orders, err := RunQueryCustomerNames(ctx, queryCustomerNames)
	if err != nil {
		return nil, err
	}

	sugar.Info("calculating unique customer count")
	customerSet := make(map[string]bool, 0)
	for _, order := range orders {
		customer := order.CustomerName
		if _, ok := customerSet[customer]; ok {
			continue
		}
		customerSet[customer] = true
	}

	sugar.Info("creating TotalSalesRevenueResponse for gRPC method GetTotalSalesRevenue")
	response := pb.CustomerCountResponse{
		Count: int64(len(customerSet)),
	}

	return &response, nil
}

// GetMerchantCount implements the GetMerchantCount method of the SalesAdminServiceServer
func (s *server) GetMerchantCount(ctx context.Context, in *pb.MerchantCountRequest) (*pb.MerchantCountResponse, error) {
	// configure logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("calling RunQueryMerchantNames method")
	orders, err := RunQueryMerchantNames(ctx, queryMerchantNames)
	if err != nil {
		return nil, err
	}

	sugar.Info("calculating unique merchant count")
	merchantSet := make(map[string]bool, 0)
	for _, order := range orders {
		merchant := order.MerchantName
		if _, ok := merchantSet[merchant]; ok {
			continue
		}
		merchantSet[merchant] = true
	}

	sugar.Info("creating TotalSalesRevenueResponse for gRPC method GetTotalSalesRevenue")
	response := pb.MerchantCountResponse{
		Count: int64(len(merchantSet)),
	}

	return &response, nil
}
