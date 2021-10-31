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

const (
	orderBatchSize = 3
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

// Client stream pattern
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

// Bi-direction stream pattern
func (s *OrderManagementServer) ProcessOrders(stream ec.OrderManagement_ProcessOrdersServer) error {
	batchMarker := 1
	var combinedShipmentMap = make(map[string]ec.CombinedShipment)

	for {
		orderId, err := stream.Recv()
		log.Printf("Reading Proc order: %s", orderId)

		if err == io.EOF {
			log.Printf("EOF: %s", orderId)
			for _, shipment := range combinedShipmentMap {
				if err := stream.Send(&shipment); err != nil {
					return err
				}
			}
			return nil
		}

		if err != nil {
			log.Println(err)
			return err
		}

		ord := s.orders[orderId.GetValue()]
		dest := ord.Destination

		shipment, found := combinedShipmentMap[dest]
		if found {
			shipment.OrdersList = append(shipment.OrdersList, &ord)
			combinedShipmentMap[dest] = shipment
		} else {
			comShip := ec.CombinedShipment{Id: "cmb - " + (dest), Status: "Processed!"}
			comShip.OrdersList = append(shipment.OrdersList, &ord) // append(comShip.OrderList, &ord)
			combinedShipmentMap[dest] = comShip
			log.Print(len(comShip.OrdersList), comShip.GetId())
		}

		if batchMarker == orderBatchSize {
			for _, comb := range combinedShipmentMap {
				log.Printf("Shipping: %v -> %v", comb.Id, len(comb.OrdersList))
				if err := stream.Send(&comb); err != nil {
					return err
				}
			}
			batchMarker = 0
			combinedShipmentMap = make(map[string]ec.CombinedShipment)
		} else {
			batchMarker++
		}
	}
}
