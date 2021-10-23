package impl

import (
	"context"
	ec "order/internal/grpc/ecommerce"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Unary communication from client to server, demostrate type of connection
// ...
type OrderServer struct {
	orders map[string]*ec.Order
}

func NewOrderServer() *OrderServer {
	o1, _ := genOrder("1")
	o2, _ := genOrder("2")

	orders := make(map[string]*ec.Order)
	orders[o1.Id] = o1
	orders[o2.Id] = o2

	return &OrderServer{orders: orders}
}

func (o *OrderServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*ec.Order, error) {
	order, exist := o.orders[orderId.Value]
	if exist {
		return order, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist.", orderId.Value)
}

func genOrder(id string) (*ec.Order, error) {
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
	return &order, nil
}
