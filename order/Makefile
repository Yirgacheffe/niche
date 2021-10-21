gen:
	@echo ''
	protoc -I . --go_out=.. ./internal/proto-files/domain/product.proto
	protoc -I . --go_out=plugins=grpc:.. ./internal/proto-files/service/product_service.proto

clean:
	@echo 'run clean tasks'

run:
	go run ./main -port 50051

build:
	go build ./cmd/grpc/server/main.go
