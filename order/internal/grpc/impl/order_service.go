package impl

import (
	"context"
	"fmt"
	"io"
	"log"
	ec "order/internal/grpc/ecommerce"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Unary communication from client to server, demostrate type of connection
// ...
type OrderManagementServer struct {
	orders map[string]ec.Order
	ec.UnimplementedOrderManagementServer
}

func NewOrderManagementServer(orders map[string]ec.Order) *OrderManagementServer {
	return &OrderManagementServer{orders: orders}
}

// Client unary pattern
func (s *OrderManagementServer) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*ec.Order, error) {
	order, exist := s.orders[orderId.Value]
	if exist {
		return &order, status.New(codes.OK, "").Err()
	}

	return nil, status.Errorf(codes.NotFound, "Order does not exist.", orderId.Value)
}

// Server stream pattern
func (s *OrderManagementServer) SearchOrders(q *wrappers.StringValue, stream ec.OrderManagement_SearchOrdersServer) error {

	for key, order := range s.orders {
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

func (s *OrderManagementServer) UpdateOrders(stream ec.OrderManagement_UpdateOrdersServer) error {

	ordersStr := "Updated Order IDs : "

	for {
		order, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&wrappers.StringValue{Value: "Orders processed " + ordersStr},
			)
		}

		s.orders[order.Id] = *order
		log.Printf("Order ID : %s - %s", order.Id, "Updated")
		ordersStr += order.Id + ", "
	}

}
