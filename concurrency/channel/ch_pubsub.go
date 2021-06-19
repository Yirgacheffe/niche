package main

import (
	"fmt"
	"sync"
)

type Pubsub struct {
	mu     sync.RWMutex
	subs   map[string][]chan string // topic -> channels
	closed bool
}

func NewPubSub() *Pubsub {
	ps := &Pubsub{}
	ps.subs = make(map[string][]chan string)
	return ps
}

// Version 1
func (ps *Pubsub) Subscribe(topic string, ch chan string) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subs[topic] = append(ps.subs[topic], ch)
}

func (ps *Pubsub) Publish(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		ch <- msg
	}
}

func (ps *Pubsub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if ps.closed {
		return
	}

	ps.closed = true

	for _, subs := range ps.subs {
		for _, ch := range subs {
			close(ch)
		}
	}
}

// Version 2
// provide a buffered channel to subscriber
// and publish message in seperate goroutine
func (ps *Pubsub) PublishV2(topic string, msg string) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return
	}

	for _, ch := range ps.subs[topic] {
		go func(ch chan string) {
			ch <- msg
		}(ch)
	}
}

func (ps *Pubsub) SubscribeV2(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 1)
	ps.subs[topic] = append(ps.subs[topic], ch)
	return ch
}

func main() {
	pubsub := NewPubSub()

	ch1 := pubsub.SubscribeV2("topic1")
	ch2 := pubsub.SubscribeV2("topic2")
	ch3 := pubsub.SubscribeV2("topic1")

	pubsub.PublishV2("topic1", "hello topic1")
	pubsub.PublishV2("topic2", "hello topic2")

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)

	pubsub.Close()

	// not send as pubsub closed
	pubsub.Publish("topic2", "hello topic2 again")
}
