package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

const listenAddr = "localhost:9010"

var indexTemplate = template.Must(template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
<script type="javascript">
	var sock = new WebSocket("ws://{{.}}/socket");
	sock.onmessage = function(m) { console.log("Received: " + m.data); }
	sock.onclose   = function(m) { console.log("Close: "    + m.code); }
	sock.send("Hello!\n");
</script>
</head>
<body></body>
</html>
`))

var partner = make(chan io.ReadWriteCloser)

// embeded io.ReadWriter instead of 'conn *websocket.Conn' to avoid Read/Write function
type socket struct {
	io.ReadWriter
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

/*
type socket struct {
    conn *websocket.Conn
    done chan bool
}

func (s socket) Read(b []byte) (int, error)  { return s.conn.Read(b) }
func (s socket) Write(b []byte) (int, error) { return s.conn.Write(b) }
*/

// Index handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	indexTemplate.Execute(w, listenAddr) // pass listenAddr as data
}

func socketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)}
	go match(s)
	<-s.done
}

func generalSockethandler(ws *websocket.Conn) {
	var s string
	fmt.Fscan(ws, &s)
	fmt.Println("Received:", s)

	fmt.Fprint(ws, "How do you do?") // write back message
}

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

	http.HandleFunc("/", indexHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))

	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Mixed http handler and websocket, demo purpose, confused with websocket.

}
