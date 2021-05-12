package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	SERVER := "localhost" + ":" + os.Args[1]

	addr, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := make([]byte, 1024)

	for {
		n, err := c.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			c.Close()
			return
		}

		fmt.Print("> ", string(buffer[0:n-1]))

		if _, err = c.Write(buffer); err != nil {
			fmt.Println(err)
			return
		}
	}

}
