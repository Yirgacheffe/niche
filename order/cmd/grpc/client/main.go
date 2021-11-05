package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "order/internal/grpc/domain"
	ec "order/internal/grpc/ecommerce"
	sv "order/internal/grpc/service"
)

const addr = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC: %v", err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	log.Println("--------------------------------------------------------------------")
	client := sv.NewProductServiceClient(conn)
	iphone13 := pb.Product{
		Name:        "iPhone 13",
		Description: "A new Iphone model to the market.",
		Price:       7899.00,
	}

	var header, trailer metadata.MD
	log.Println("Simple test: ")

	if id, err := client.AddProduct(ctx, &iphone13, grpc.Header(&header), grpc.Trailer(&trailer)); err != nil {
		log.Fatalf("not able to add product: %v", err)
	} else {
		prod, err := client.GetProduct(ctx, id)
		if err != nil {
			log.Fatalf("retrieve product failed: %v", err)
		}

		prodJSON, _ := json.Marshal(prod)
		log.Printf("Get product model in details: %#v", string(prodJSON))
	}

	log.Println("--------------------------------------------------------------------")
	log.Println("Communication test: ")

	cli := ec.NewOrderManagementClient(conn)
	if order, err := cli.GetOrder(ctx, &wrappers.StringValue{Value: "102"}); err != nil {
		errStat := status.Convert(err)

		log.Printf("%d: %s\n", errStat.Code(), errStat.Message())
		log.Fatalf("not able to get order: %v", err)
	} else {
		log.Printf("Get order with id: %s, %v", "1", order)
	}

	log.Println("--------------------------------------------------------------------")
	log.Println("Server stream: ")
	// Search: server stream
	searchStream, _ := cli.SearchOrders(ctx, &wrappers.StringValue{Value: "iPhone"})
	for {
		order, err := searchStream.Recv()
		if err == io.EOF {
			break
		}

		log.Println("Search result: ", order)
	}

	log.Println("--------------------------------------------------------------------")
	log.Println("Client stream: ")
	// Update: client stream
	updOrder1 := ec.Order{Id: "102", Items: []string{"Google Pixel 3A", "Google Pixel Book"}, Destination: "San Jose, UK", Price: 2800.00}
	updOrder2 := ec.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination: "Zhi hui shan No.3", Price: 3102.39}
	updOrder3 := ec.Order{Id: "105", Items: []string{"Apple Watch S4"}, Destination: "Wood View TX", Price: 1100.00}

	updStream, err := cli.UpdateOrders(ctx)
	if err != nil {
		log.Fatalf("%v.UpdateOrder(_) = _, %v", cli, err)
	}

	if err := updStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updStream, updOrder1, err)
	}

	if err := updStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updStream, updOrder2, err)
	}

	if err := updStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updStream, updOrder3, err)
	}

	updateRes, err := updStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updStream, err, nil)
	}

	log.Printf("Update Order Res: %s", updateRes)

	log.Println("--------------------------------------------------------------------")
	log.Println("Bi-directional stream: ")

	procOrder, err := cli.ProcessOrders(ctx)
	if err != nil {
		log.Fatalf("%v.ProcessOrders(_) = _, %v", cli, err)
	}

	if err := procOrder.Send(&wrappers.StringValue{Value: "102"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", cli, "102", err)
	}

	if err := procOrder.Send(&wrappers.StringValue{Value: "104"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", cli, "104", err)
	}

	if err := procOrder.Send(&wrappers.StringValue{Value: "105"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", cli, "105", err)
	}

	//....
	ch := make(chan struct{})
	go asyncClientBidirectionalRPC(procOrder, ch)
	time.Sleep(time.Millisecond * 1000)

	if err := procOrder.Send(&wrappers.StringValue{Value: "101"}); err != nil {
		log.Fatalf("%v.Send(%v) = %v", cli, "101", err)
	}

	if err := procOrder.CloseSend(); err != nil {
		log.Fatal(err)
	}

	ch <- struct{}{}

	log.Println("--------------------------------------------------------------------")
	log.Println("Bi-directional stream: ")
	// dry-run command test
	ctxx := metadata.NewOutgoingContext(ctx, metadata.Pairs("dry-run", "1"))
	did, err := client.AddProduct(ctxx, &iphone13)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("dry-run: get id: " + did)
	log.Println("--------------------------------------------------------------------")

}

func asyncClientBidirectionalRPC(streamProcOrder ec.OrderManagement_ProcessOrdersClient, c chan struct{}) {
	for {
		combinedShipmet, err := streamProcOrder.Recv()
		if err == io.EOF {
			break
		}

		log.Println("Combined shipment: ", combinedShipmet.OrdersList)
	}
	<-c
}
