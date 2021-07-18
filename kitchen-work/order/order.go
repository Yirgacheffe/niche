package order

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Order struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Time      int    `json:"prepTime"`
	CourierId int
	StartTime int64
}

type OrderManager struct {
	CourierNum int
}

func NewOrderManager(courierNum int) (*OrderManager, error) {
	if courierNum < 1 {
		return nil, errors.New("Invalid courier number.")
	}

	return &OrderManager{CourierNum: courierNum}, nil
}

func (m *OrderManager) Cooking(done <-chan bool, orders []Order) <-chan Order {
	readyOrders := make(chan Order)

	go func() {
		defer close(readyOrders)
		for _, o := range orders {
			fmt.Printf("Order [%s]: Preparing in %d seconds...\n", o.Id, o.Time)
			cookTime := time.Duration(o.Time) * time.Second
			o.StartTime = time.Now().UnixNano() / int64(time.Millisecond)

			select {
			case <-done:
				return
			case <-time.After(cookTime):
				readyOrders <- o
			}
		}
	}()

	return readyOrders
}

func (m *OrderManager) Receive(done <-chan bool, orders <-chan Order) <-chan Order {
	receivedOrders := make(chan Order)
	randFn := func() int { return rand.Intn(m.CourierNum) + 1 }

	go func() {
		defer close(receivedOrders)
		for {
			select {
			case <-done:
				return
			case o, ok := <-orders:
				if ok {
					o.CourierId = randFn()
					fmt.Printf("Received: %s, assign courier: [%d]\n", o.Id, o.CourierId)
					receivedOrders <- o
				} else {
					return
				}
			}
		}
	}()
	return receivedOrders
}
