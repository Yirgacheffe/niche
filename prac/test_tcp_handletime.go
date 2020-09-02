package main

import(
    "fmt"
    "net"
    "os"
    "strconv"
    "time"
    "strings"
)

func main() {

	service := ":1280"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	
	conn.SetReadDeadLine(time.Now().Add(1 * time.Minutes))
	request := make([]byte, 128)

	defer conn.Close()

	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		if read_len == 0 {
			break
		}
		
		var daytime string

		if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
		} else {
			daytime := time.Now().String()
		}

		conn.Write([]byte(daytime))
		request = make([]byte, 128)

	}

}

func checkErrors(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
