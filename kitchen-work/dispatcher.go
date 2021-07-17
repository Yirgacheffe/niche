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

const (
	courierNum  = 3
	rawDataFile = "./data/input.json"
)

func DispatchFIFO() {
	fmt.Println("------------------------- START..... --------------------------")
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

	readyOrder := make(chan order.Order)
	defer close(readyOrder)

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
	for _, v := range cMgr.Couriers {
		go v.NotifyToPickup(done, readyOrder)
	}

	// Preparing orders, couriers waitting for picking up the order
	for o := range oMgr.Cooking(done, needCooked) {
		readyOrder <- o
	}
	fmt.Println("------------------------ FINISH..... --------------------------")
}

func DispatchMatched() {
	fmt.Println("------------------------- START..... --------------------------")

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

	readyOrder := make(chan order.Order)
	defer close(readyOrder)

	rawOrders, err := loadOrders(rawDataFile)
	if err != nil {
		fmt.Println("Load data error: ", err)
		return
	}

	// Received orders, assign to couriers
	received := oMgr.Receive(done, generate(done, rawOrders))
	var needCooked []order.Order

	for o := range received {
		needCooked = append(needCooked, o) // TODO: Split the order and send to courier
	}

	// Get couriers ready, arrival to kitchen...
	for _, v := range cMgr.Couriers {
		go v.NotifyToPickup(done, readyOrder)
	}

	// Preparing orders, couriers waitting for picking up the order
	for o := range oMgr.Cooking(done, needCooked) {
		readyOrder <- o
	}

	fmt.Println("------------------------ FINISH..... --------------------------")
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
		workGen := time.Tick(500 * time.Millisecond)

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
