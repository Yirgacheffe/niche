protoc -I . --go_out=. ./bitbucket-repos/internal/proto-files/domain/repository.proto
protoc -I . --go_out=plugins=grpc:. ./bitbucket-repos/internal/proto-files/service/repository-service.proto