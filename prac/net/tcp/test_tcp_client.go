package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide host:port")
		return
	}

	CONNECT := os.Args[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n') // ignore error right now
		fmt.Fprintf(c, text+"\n")

		if strings.TrimSpace((string(text))) == "STOP" {
			fmt.Println("TCP client exiting ...")
			return
		}

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
	}

}
