package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const listenAddr = "localhost:9010"

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Waiting for a partner...\n")
	select {
	case partner <- c:
		// Now handled by other goroutine
		// current goroutine will exit
	case p := <-partner:
		chat(p, c) // Got a partner, start to chat
	}
}

func chat(p, c io.ReadWriteCloser) {
	fmt.Fprint(p, "Found one! Say hi.")
	fmt.Fprint(c, "Found one! Say hi.")

	errc := make(chan error, 1)
	go cp(p, c, errc)
	go cp(c, p, errc)

	if err := <-errc; err != nil {
		log.Println(err)
	}

	p.Close()
	c.Close()
}

func cp(w io.ReadWriteCloser, r io.ReadWriteCloser, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err // err could be nil, length is 1, no blocking
}

func main() {

	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go match(c)
	}

}
