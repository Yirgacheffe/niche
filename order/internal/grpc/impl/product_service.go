package impl

import (
	"context"
	pb "order/internal/grpc/domain"
	sv "order/internal/grpc/service"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type ProductServer struct {
	productMap map[string]*pb.Product
	sv.UnimplementedProductServiceServer
}

func NewProductServer() *ProductServer {
	return &ProductServer{}
}

func (s *ProductServer) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {

	if dryRun(ctx) {
		return &pb.ProductID{
			Value: "123e4567-e89b-12d3-a456-426655440000"}, nil
	}

	id, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, "Error while generate Product ID", err)
	}

	in.Id = id.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}

	s.productMap[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *ProductServer) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	v, exists := s.productMap[in.Value]
	if exists {
		return v, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Product does not exist.", in.Value)
}

func dryRun(ctx context.Context) bool {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return false
	}

	val, ok := md["dry-run"]
	if !ok {
		return false
	}

	if len(val) < 1 {
		return false
	}

	return val[0] == "1" // dry-run 1: true others: false
}
