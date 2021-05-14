package main

import (
	"fmt"
	"net"
)

func main() {

	netAddr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return
	}

	c, err := net.ListenIP("ip4:icmp", netAddr)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := make([]byte, 1024)
	n, _, err := c.ReadFrom(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("% X\n", buf[0:n]) // need security privilege

}
