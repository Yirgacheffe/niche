package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

var chain = NewChain(2)

// Prefix is a Markov chain prefix of one or more words.
type Prefix []string

func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Chain contains a map ("chain") of prefixes to
// a list of suffixes.
type Chain struct {
	chain     map[string][]string
	prefixLen int
}

func NewChain(prefixLen int) *Chain {
	return &Chain{make(map[string][]string), prefixLen}
}

func (c *Chain) Write(b []byte) (int, error) {
	// Convert []byte to prefix and suffix into chain.
	key := string(b)

	p := make(Prefix, c.prefixLen)
	c.chain[key] = append(c.chain[key], s)
	p.Shift(s)
	return len(b), nil
}

func (c *Chain) Generate(n int) string {
	p := make(Prefix, c.prefixLen)
	var words []string

	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}

		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}

	return strings.Join(words, " ")
}

// Bot return an io.ReadWriteCloser the responds to incoming
// write with a generated sentence.
func Bot() io.ReadWriteCloser {
	r, out := io.Pipe()
	return bot{r, out}
}

type bot struct {
	io.ReadCloser
	out io.Writer
}

func (b bot) Write(buf []byte) (int, error) {
	go b.speak()
	return len(buf), nil
}

func (b bot) speak() {
	time.Sleep(time.Second)
	b.out.Write([]byte(chain.Generate(10)))
}

const listenAddr = "localhost:9010"

type socket struct {
	io.Reader
	io.Writer
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

// SocketHandler handle http connection then start to match and chat
func socketHandler(ws *websocket.Conn) {
	r, w := io.Pipe()
	go func() {
		_, err := io.Copy(io.MultiWriter(w, chain), ws)
		w.CloseWithError(err)
	}()

	s := socket{r, w, make(chan bool)}
	go match(s)

	<-s.done
}

// Match a partner then start the chat
var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Waiting for a partner...\n")
	select {
	case partner <- c:
		// Now handled by other goroutine
		// current goroutine will exit
	case p := <-partner:
		chat(p, c) // Got a partner, start to chat
	case <-time.After(5 * time.Second):
		chat(Bot(), c)
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

	http.Handle("/socket", websocket.Handler(socketHandler))

	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		fmt.Println(err)
		return
	}

}
