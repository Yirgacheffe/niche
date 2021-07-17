package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	courierId int
	time      int
	name      string
}

type Courier struct {
	id       int
	receiver chan Order
}

const maxCourierNum = 3

var orders []Order = []Order{
	Order{1, 1, 8, "courier 1"}, Order{2, 2, 4, "courier 2"}, Order{3, 3, 2, "courier 3"},
	Order{4, 1, 5, "courier 1"}, Order{5, 2, 6, "courier 2"}, Order{6, 3, 9, "courier 3"},
	Order{7, 1, 6, "courier 1"}, Order{8, 2, 6, "courier 2"}, Order{9, 3, 6, "courier 3"},
}

func split(orders []Order) map[int][]Order {
	m := make(map[int][]Order)

	for _, o := range orders {
		id := o.courierId

		if v, ok := m[id]; !ok {
			m[id] = []Order{o}
		} else {
			m[id] = append(v, o)
		}
	}
	return m // -------------------------!
}

func doWork(orders []Order) {
	for _, v := range orders {
		fmt.Println(v)
	}
}

func main() {
	/*
		for _, v := range split(orders) {
			go doWork(v)
		}
	*/
	done := make(chan bool)
	defer close(done)

	k := make(map[int]chan string)
	k[1] = make(chan string)
	k[2] = make(chan string)

	go func() {
		k[1] <- "a"
	}()

	go func() {
		k[2] <- "b"
	}()

	for i, v := range k {
		fmt.Println(i)
		fmt.Println(<-v)
	}

	//
	xyz := cookingSplit(done, orders)

	for k, v := range xyz {

		print := func(k int, orders <-chan Order) {
			fmt.Println(k)

			for o := range orders {
				fmt.Println(o)
			}
		}

		go print(k, v)

	}

	time.Sleep(200 * time.Second)
}

/*

	key := o.courierId
	if v, ok := courierOrder[key]; !ok {
		courierOrder[key] = []Order{o}
	} else {
		courierOrder[key] = append(v, o)
	}
*/

func cookingSplit(done <-chan bool, orders []Order) map[int]chan Order {
	readyChs := make(map[int](chan Order))

	for i := 1; i <= 3; i++ {
		readyChs[i] = make(chan Order)
	}

	go func() {
		for _, v := range readyChs {
			defer close(v)
		}

		for _, o := range orders {
			cookTime := time.Duration(o.time) * time.Second
			select {
			case <-done:
				return
			case <-time.After(cookTime):
				readyChs[o.courierId] <- o
			}
		}
	}()

	return readyChs
}
