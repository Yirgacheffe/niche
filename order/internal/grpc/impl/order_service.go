package impl

import (
	"context"
	"fmt"
	"log"
	ec "order/internal/grpc/ecommerce"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Unary communication from client to server, demostrate type of connection
// ...
type OrderServer struct {
	orders map[string]ec.Order
}

func NewOrderServer() *OrderServer {
	o1, _ := genOrder("1")
	o2, _ := genOrder("2")
	o3, _ := genOrder("3")

	orders := make(map[string]ec.Order)
	orders[o1.Id] = o1
	orders[o2.Id] = o2
	orders[o3.Id] = o3

	return &OrderServer{orders: orders}
}

// Client unary pattern
func (o *OrderServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*ec.Order, error) {
	order, exist := o.orders[orderId.Value]
	if exist {
		return &order, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist.", orderId.Value)
}

// Server stream pattern
func (o *OrderServer) SearchOrders(q *wrappers.StringValue, stream ec.OrderManagement_SearchOrdersServer) error {

	for key, order := range o.orders {
		log.Print(key, order)

		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, q.Value) {
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("error sending order to stream: %v", err)
				}

				log.Print("Match Order Found: " + key)
				break
			}
		}
	}

	return nil
}

func genOrder(id string) (ec.Order, error) {
	/*
		id, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
	*/
	order := ec.Order{
		Id:          id,
		Items:       []string{"iphone13", "distributed system design", "Yamazki"},
		Description: "Multi-books, mobild phone",
		Price:       255.60,
		Destination: "Hua Yuan Chan Ye Yuan No.350",
	}
	return order, nil
}
