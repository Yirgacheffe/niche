package courier

import (
	"errors"
	"fmt"
	"kitchen-work/order"
	"math/rand"
	"time"
)

type Courier struct {
	Id int
}

type CourierManager struct {
	Couriers []*Courier
}

func NewCourierManager(maxNbr int) (*CourierManager, error) {
	if maxNbr < 1 {
		return nil, errors.New("Invalid max number.")
	}

	var couriers []*Courier
	for i := 1; i <= maxNbr; i++ {
		couriers = append(couriers, newCourier(i))
	}

	return &CourierManager{Couriers: couriers}, nil
}

func newCourier(Id int) *Courier {
	return &Courier{Id: Id}
}

func (c *Courier) NotifyToPickup(done <-chan bool, orders <-chan order.Order) {

	fmt.Printf("Courier [%d]: notified by kitchen.\n", c.Id)
	randFn := func() int {
		return rand.Intn(12) + 3
	}()

	arrivalDuration := time.Duration(randFn) * time.Second
	fmt.Printf("Courier [%d]: is approching..., remain %d seconds\n", c.Id, randFn)

	// Wait for courier arrival
	select {
	case <-done:
		return
	case <-time.After(arrivalDuration):
		fmt.Printf("Courier [%d]: arrived\n", c.Id)
	}

	// Pickup orders
	go func() {
		for {
			select {
			case <-done:
				return
			case o, ok := <-orders:
				if ok {
					fmt.Printf("Courier [%d]: pickup order [%s] \n", c.Id, o.Id)
				} else {
					return
				}
			}
		}
	}()

}
