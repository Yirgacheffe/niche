package main

import (
	"bitbucket-repos/internal/grpc/domain"
	"bitbucket-repos/internal/grpc/service"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {

	srvAddress := "localhost:7000"
	conn, err := grpc.Dial(srvAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := service.NewRepositoryServiceClient(conn)

	for i := range [10]int{} {
		repositoryModel := domain.Repository{
			Id:        int64(i),
			IsPrivate: true,
			Name:      string("Grpc-Demo"),
			UserId:    1245,
		}

		if repMsg, err := client.Add(context.Background(), &repositoryModel); err != nil {
			panic(fmt.Sprintf("Was not able to insert Record: %v", e))
		} else {
			fmt.Println("Record Inserted..")
			fmt.Println(repMsg)
			fmt.Println("========================")
		}
	}

}
