package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + os.Args[1]
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()
	rand.Seed(time.Now().Unix())
	buffer := make([]byte, 1024)

	for {
		n, addr, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Output the text got from client
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		replyData := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(replyData))

		_, err = c.WriteToUDP(replyData, addr)
		if err != nil {
			fmt.Println("Send data back error: ", err)
			return
		}
	}

}
