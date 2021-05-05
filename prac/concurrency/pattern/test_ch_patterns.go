package main

import "fmt"

func main() {

	orDone := func(done <-chan bool, c <-chan int) <-chan int {
		valStream := make(chan int)

		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	tee := func(done <-chan bool, in <-chan int) (_, _ <-chan int) {
		out1 := make(chan int)
		out2 := make(chan int)

		go func() {
			defer close(out1)
			defer close(out2)
			for v := range orDone(done, in) {
				var out1, out2 = out1, out2
				for i := 0; i < 2; i++ {
					select {
					case <-done:
						return
					case out1 <- v:
						out1 = nil // set this to nil blocked the op, so out2 have chance to write
					case out2 <- v:
						out2 = nil // set this to nil blocked the op, so out1 have chance to write
					}
				}
			}
		}()

		return out1, out2
	}

	bridge := func(done <-chan bool, chanStream <-chan <-chan int) <-chan int {
		valStream := make(chan int)

		go func() {
			defer close(valStream)

			for {
				var stream <-chan int
				select {
				case <-done:
					return
				case maybeStream, ok := <-chanStream:
					if ok == false {
						return
					}
					stream = maybeStream
				}

				for v := range orDone(done, stream) {
					select {
					case <-done:
					case valStream <- v:
					}
				}
			}
		}()

		return valStream
	}

	genVals := func() <-chan <-chan int {
		chanStream := make(chan (<-chan int))

		go func() {
			defer close(chanStream)
			for i := 0; i < 10; i++ {
				stream := make(chan int, 1)
				stream <- i
				close(stream)
				chanStream <- stream
			}
		}()

		return chanStream
	}

	done := make(chan bool)
	defer close(done)

	for v := range bridge(done, genVals()) {
		fmt.Printf("%v ", v)
	}

}
