gen:
	@echo ''
	protoc -I . --go_out=:..                 ./internal/proto-files/domain/product.proto
	protoc -I . --go-grpc_out=..             ./internal/proto-files/service/product_service.proto
	protoc -I . --go_out=.. --go-grpc_out=.. ./internal/proto-files/ecommerce/order_management.proto
	

clean:
	@echo 'run clean tasks'

run:
	@echo 'run gRPC server...'
	./server -p 50051

build:
	@echo 'build server...'
	go build -o server ./cmd/grpc/server/*.go

client:
	go build -o client ./cmd/grpc/client/*.go

test:
	@echo 'run client for test...'
	./client

purge:
	rm -rf ./client
	rm -rf ./server