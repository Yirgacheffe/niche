package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"os"
	"sharedRPC"
)

type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (t *MyInterface) Multiply(args *sharedRPC.MyFloats, reply *float64) error {
	*reply = args.A1 * args.A2
	return nil
}

func (t *MyInterface) Power(args *sharedRPC.MyFloats, reply *float64) error {
	*reply = Power(args.A1, args.A2)
	return nil
}

func main() {

	PORT := ":19010"
	if len(os.Args) != 1 {
		PORT = ":" + os.Args[1]
	}

	myInterface := new(MyInterface) // create a new instance
	rpc.Register(myInterface)

	t, err := net.ResolveTCPAddr("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}

}
