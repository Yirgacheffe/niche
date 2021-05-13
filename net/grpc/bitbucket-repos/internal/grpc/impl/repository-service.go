package impl

import (
	"bitbucket-repos/internal/grpc/domain"
	"bitbucket-repos/internal/grpc/service"
	"context"
	"log"
	"strconv"
	// "../domain"
	// "../service"
)

//RepositoryServiceGrpcImpl is a implementation of RepositoryService Grpc Service.
type RepositoryServiceGrpcImpl struct {
}

//NewRepositoryServiceGrpcImpl returns the pointer to the implementation.
func NewRepositoryServiceGrpcImpl() *RepositoryServiceGrpcImpl {
	return &RepositoryServiceGrpcImpl{}
}

// Add function implementation of gRPC service.
func (serviceImpl *RepositoryServiceGrpcImpl) Add(ctx context.Context, in *domain.Repository) (*service.AddRepositoryRresponse, error) {

	log.Println("Received request for adding repository with id " + strconv.FormatInt(in.Id, 10))
	log.Println("Repository persisted to the storage")

	return &service.AddRepositoryResponse{
		AddedRepository: in,
		Error:           nil,
	}, nil

}
