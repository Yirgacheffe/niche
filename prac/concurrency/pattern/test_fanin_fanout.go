package main

import (
	"fmt"
	"sync"
)

func getInputChan() <-chan int {

	ch := make(chan int, 100)
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for n := range numbers {
			ch <- n
		}
		close(ch)
	}()

	return ch // return input channel

}

func getSquareChan(input <-chan int) <-chan int {

	out := make(chan int, 100)

	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()

	return out // output square channel

}

// fanin
func merge(sqrChan ...<-chan int) <-chan int {

	var wg sync.WaitGroup
	merged := make(chan int, 100)

	wg.Add(len(sqrChan))

	output := func(sc <-chan int) {
		for sqr := range sc {
			merged <- sqr
		}
		wg.Done()
	}

	for _, optCh := range sqrChan {
		go output(optCh)
	}

	// another go routine wait for all
	// result finished
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged // return merged channel

}

func main() {

	fmt.Println("[main] main started......")
	fmt.Println("[main] Calc: Input -> Square * 2 -> Merged")

	chInputNums := getInputChan()

	chOptSquare1 := getSquareChan(chInputNums)
	chOptSquare2 := getSquareChan(chInputNums)

	chMergedSqr := merge(chOptSquare1, chOptSquare2)

	sqrSum := 0
	for sqr := range chMergedSqr {
		sqrSum += sqr
	}

	fmt.Println("Sum of squares between 0-9 is:", sqrSum)

}
