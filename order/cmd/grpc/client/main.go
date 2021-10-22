package main

import (
	"context"
	"encoding/json"
	"log"

	"google.golang.org/grpc"

	pb "order/internal/grpc/domain"
	sv "order/internal/grpc/service"
)

const addr = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC: %v", err)
	}

	defer conn.Close()

	client := sv.NewProductServiceClient(conn)
	iphone13 := pb.Product{
		Name:        "iPhone 13",
		Description: "A new Iphone model to the market.",
		Price:       7899.00,
	}

	if id, err := client.AddProduct(context.Background(), &iphone13); err != nil {
		log.Fatalf("not able to add product: %v", err)
	} else {
		prod, err := client.GetProduct(context.Background(), id)
		if err != nil {
			log.Fatalf("retrieve product failed: %v", err)
		}

		prodJSON, _ := json.Marshal(prod)
		log.Printf("Get product model in details: %#v", string(prodJSON))
	}

}
