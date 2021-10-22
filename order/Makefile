gen:
	@echo ''
	protoc -I . --go_out=.. ./internal/proto-files/domain/product.proto
	protoc -I . --go_out=plugins=grpc:.. ./internal/proto-files/service/product_service.proto

clean:
	@echo 'run clean tasks'

run:
	@echo 'run gRPC server...'
	./server -port 50051

build:
	@echo 'build server...'
	go build -o server ./cmd/grpc/server/main.go

client:
	go build -o client ./cmd/grpc/client/main.go

test:
	@echo 'run client for test...'
	./client

purge:
	rm -rf ./client
