package main

import (
	"fmt"

	"github.com/trustmaster/goflow"
)

// Demo from goflow, go details later
type Greeter struct {
	Name <-chan string // input port
	Res  chan<- string // output port
}

func (c *Greeter) Process() {
	for name := range c.Name {
		greeting := fmt.Sprintf("Hello, %s!", name)
		c.Res <- greeting
	}
}

type Printer struct {
	Line <-chan string // inport
}

func (c *Printer) Process() {
	for line := range c.Line {
		fmt.Println(line)
	}
}

func NewGreetingApp() *goflow.Graph {
	n := goflow.NewGraph()

	n.Add("greeter", new(Greeter))
	n.Add("printer", new(Printer))
	n.Connect("greeter", "Res", "printer", "Line")
	n.MapInPort("In", "greeter", "Name")
	return n
}

func main() {
	net := NewGreetingApp()

	in := make(chan string)
	net.SetInPort("In", in)

	wait := goflow.Run(net)
	in <- "John"
	in <- "Boris"
	in <- "Hanna"

	close(in)
	<-wait
}
