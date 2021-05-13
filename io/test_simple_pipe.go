package main

import (
	"fmt"
	"io"
)

func main() {

	src, dst := io.Pipe()

	// start goroutine writes data to 'dst'
	go func() {
		dst.Write([]byte("DATA_1"))
		dst.Write([]byte("DATA_2"))
		dst.Close() // indicate EOF
	}()

	packet := make([]byte, 6)

	bytesRead1, err1 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead1, packet, err1)

	bytesRead2, err2 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead2, packet, err2)

	bytesRead3, err3 := src.Read(packet)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead3, packet, err3)

}
