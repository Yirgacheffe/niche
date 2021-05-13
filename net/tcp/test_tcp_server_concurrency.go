package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func f(n int) int {
	fn := make(map[int]int)

	for i := 0; i <= n; i++ {
		var t int
		if i == 0 || i == 1 {
			t = i
		} else {
			t = fn[i-1] + fn[i-2]
		}
		fn[i] = t
	}

	return fn[n]
}

func handleConn(c net.Conn) {

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}

		// error happened, return this
		fibo := "-1\n"

		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}

		c.Write([]byte(string(fibo)))
	}

	time.Sleep(3 * time.Second)
	c.Close() // Close conn once got 'STOP' signal from client

}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + os.Args[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConn(c) // handle each client connection
	}

}
