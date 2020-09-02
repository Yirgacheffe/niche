package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {

	var err error

	for {
		var reply string
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive...")
			break
		}

		msg := "Received: " + reply
		fmt.Println("Received back from client: " + reply)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send...")
			break
		}
	}

}

func main() {

	http.Handle("/", websocket.Handler(Echo))
	
	if err := http.ListenAndServe(":17620", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
