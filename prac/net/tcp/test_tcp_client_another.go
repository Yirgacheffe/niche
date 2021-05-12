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
		fmt.Println("Please provide a server:port string!")
		return
	}

	CONNECT := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		fmt.Println("DialTCP:", err.Error())
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n") // Get from stdin then send text to conn

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->:" + message)
	}

}
