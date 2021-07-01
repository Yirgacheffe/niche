package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var Password = secret{password: "myPassword"}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("L Change")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	var showFunction func(c *secret) string

	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}

	fmt.Println("Pass:", showFunction(&Password))
	var wg sync.WaitGroup
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change(&Password, "123355")
	}()

	wg.Wait()
	fmt.Println("Pass:", showFunction(&Password))
}
