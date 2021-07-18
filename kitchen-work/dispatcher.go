package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kitchen-work/courier"
	"kitchen-work/order"
	"log"
	"os"
	"time"
)

// TODO: these 2 parameters should be get from command line for improvement
const (
	courierNum  = 5
	rawDataFile = "./data/input.json"
)

func DispatchFIFO() {
	fmt.Println("----------------------- FIFO START ..... ----------------------------")
	cMgr, err := courier.NewCourierManager(courierNum)
	if err != nil {
		fmt.Println("Init error: ", err)
		return
	}
	oMgr, err := order.NewOrderManager(courierNum)
	if err != nil {
		fmt.Println("Init error: ", err)
		return
	}

	// Setup all goroutine context, close all after complete
	done := make(chan bool)
	defer close(done)

	rawOrders, err := loadOrders(rawDataFile)
	if err != nil {
		fmt.Println("Load data error: ", err)
		return
	}

	// Received orders, assign to couriers
	received := oMgr.Receive(done, generate(done, rawOrders))
	var needCooked []order.Order

	for o := range received {
		needCooked = append(needCooked, o)
	}

	// Get couriers ready, arrival to kitchen...
	readyOrder := make(chan order.Order)
	defer close(readyOrder)

	for _, v := range cMgr.Couriers {
		go v.NotifyToPickup(done, readyOrder)
	}

	// Preparing orders, couriers waitting for picking up the order
	for o := range oMgr.Cooking(done, needCooked) {
		readyOrder <- o
	}

	// TODO: the statistic requirement should goes here

	fmt.Println("----------------------- FIFO FINISH ..... ---------------------------")
}

func DispatchMatched() {
	fmt.Println("----------------------- MATCHED START .... --------------------------")
	cMgr, err := courier.NewCourierManager(courierNum)
	if err != nil {
		fmt.Println("Init error: ", err)
		return
	}
	oMgr, err := order.NewOrderManager(courierNum)
	if err != nil {
		fmt.Println("Init error: ", err)
		return
	}

	// Setup all goroutine context, close all after complete
	done := make(chan bool)
	defer close(done)

	rawOrders, err := loadOrders(rawDataFile)
	if err != nil {
		fmt.Println("Load data error: ", err)
		return
	}

	// Received orders, assign to couriers
	received := oMgr.Receive(done, generate(done, rawOrders))
	var needCooked []order.Order
	for o := range received {
		needCooked = append(needCooked, o)
	}

	// initial splited match order
	splited := make(map[int]chan order.Order)
	for _, v := range cMgr.Couriers {
		ready := make(chan order.Order)
		splited[v.Id] = ready

		go func(c *courier.Courier) {
			c.NotifyToPickup(done, ready)
		}(v)
	}

	// Preparing orders, couriers waitting for picking up the order
	for o := range oMgr.Cooking(done, needCooked) {
		splited[o.CourierId] <- o
	}

	// TODO: the statistic requirement should goes here

	fmt.Println("----------------------- MATCHED FINISH ... --------------------------")
}

// loadOrders
func loadOrders(filePath string) ([]order.Order, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Open input file error:", err)
	}
	defer f.Close()

	rawJson, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var orders []order.Order
	if err = json.Unmarshal(rawJson, &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func generate(done <-chan bool, orders []order.Order) <-chan order.Order {
	genOrders := make(chan order.Order)
	go func() {
		defer close(genOrders)
		workGen := time.Tick(10 * time.Millisecond)

		for _, o := range orders {
			select {
			case <-done:
				return
			case <-workGen:
				genOrders <- o
			}
		}
	}()
	return genOrders // Generate order stream, simulate every 0.5 second...
}
