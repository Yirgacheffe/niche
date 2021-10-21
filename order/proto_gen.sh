protoc -I . --go_out=. ./internal/proto-files/domain/product.proto
protoc -I . --go_out=. ./internal/proto-files/service/product_service.proto