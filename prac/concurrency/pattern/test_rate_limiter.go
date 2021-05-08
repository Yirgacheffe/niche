package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Wrap rate.Every function
func Per(eventCnt int, duration time.Duration) rate.Limit {
	return rate.Every(duration / time.Duration(eventCnt))
}

type RateLimiter interface {
	Wait(context.Context) error
	Limit() rate.Limit
}

// Get new multi-limiter
func NewMultiLimiter(limiters ...RateLimiter) *multiLimiter {
	byLimit := func(i, j int) bool {
		return limiters[i].Limit() < limiters[j].Limit()
	}

	sort.Slice(limiters, byLimit)
	return &multiLimiter{limiters: limiters}
}

// Type of rate limiters
type multiLimiter struct {
	limiters []RateLimiter
}

func (l *multiLimiter) Wait(ctx context.Context) error {
	for _, l := range l.limiters {
		if err := l.Wait(ctx); err != nil {
			return err
		}
	}
	return nil // loop the limiter, no error happen
}

func (l *multiLimiter) Limit() rate.Limit {
	return l.limiters[0].Limit()
}

// Open Api, database and disk simulate the actual work
func Open() *APIConnection {
	return &APIConnection{
		networkLimit: NewMultiLimiter(
			rate.NewLimiter(Per(3, time.Second), 3),
		),
		diskLimit: NewMultiLimiter(
			rate.NewLimiter(rate.Limit(1), 1),
		),
		apiLimit: NewMultiLimiter(
			rate.NewLimiter(Per(2, time.Second), 2),
			rate.NewLimiter(Per(10, time.Minute), 10),
		),
	}
}

type APIConnection struct {
	networkLimit, diskLimit, apiLimit RateLimiter
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := NewMultiLimiter(a.apiLimit, a.networkLimit).Wait(ctx); err != nil {
		return err
	}

	fmt.Println("execute ResolveAddress logic here")
	return nil
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := NewMultiLimiter(a.apiLimit, a.diskLimit).Wait(ctx); err != nil {
		return err
	}

	fmt.Println("execute ReadFile logic here to simulate work")
	return nil
}

func main() {

	defer log.Printf("Done")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	apiConn := Open()

	var wg sync.WaitGroup
	wg.Add(40)

	for i := 0; i < 20; i++ {
		defer wg.Done()
		go func() {
			err := apiConn.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}

	for i := 0; i < 20; i++ {
		defer wg.Done()
		go func() {
			err := apiConn.ReadFile(context.Background())
			if err != nil {
				log.Printf("cannot ReadFile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	wg.Wait()

}
