package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	// dial 'rpc.DefaultRPCPath' endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9093")

	// Temp common student model
	var donald Student
	id := 1

	// *----------------------------

	// Get student with some id
	err := client.Call("College.Get", id, &donald)

	if err != nil {
		fmt.Println("Error: College.Get", err)
	} else {
		name := donald.FullName()
		fmt.Printf("Success: '%s' found with id %d\n", name, id)
	}

	// *----------------------------

	// Temp student for insertion
	t := Student{id, "Donald", "Trump"}

	err = client.Call("College.Add", t, &donald)
	if err != nil {
		fmt.Println("Error: College.Add", err)
	} else {
		name := donald.FullName()
		fmt.Printf("Success: '%s' created with id=%d\n", name, id)
	}

}
