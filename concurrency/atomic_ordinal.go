package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

// SafeMap uses a mutex to allow getting and setting in a thread-safe way
type SafeMap struct {
	mu *sync.RWMutex
	m  map[string]string
}

func NewSafeMap() SafeMap {
	return SafeMap{&sync.RWMutex{}, make(map[string]string)}
}

func (s *SafeMap) Set(k, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.m[k] = v
}

func (s *SafeMap) Get(k string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if v, ok := s.m[k]; ok {
		return v, nil
	}
	return "", errors.New("key not found")
}

// NewOrdinal returns ordinal with once setup
type Ordinal struct {
	once    *sync.Once
	ordinal uint64
}

func NewOrdinal() *Ordinal {
	return &Ordinal{once: &sync.Once{}}
}

func (o *Ordinal) Init(v uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, v)
	})
}

func (o *Ordinal) GetOrdinal() uint64 {
	return atomic.LoadUint64(&o.ordinal)
}

func (o *Ordinal) Increment() {
	atomic.AddUint64(&o.ordinal, 1)
}

func main() {
	o := NewOrdinal()
	s := NewSafeMap()

	o.Init(1123)
	fmt.Println("initial ordinal is:", o.GetOrdinal())

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			s.Set(fmt.Sprint(i), "success")
			o.Increment()
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		v, err := s.Get(fmt.Sprint(i))
		if err != nil || v != "success" {
			panic(err)
		}
	}

	fmt.Println("final ordinal is:", o.GetOrdinal())
	fmt.Println("all keys found and marked as: 'success'")
}
