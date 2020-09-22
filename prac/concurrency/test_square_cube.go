package main

import "fmt"

func calcSquares(nbr int, squareop chan int) {
	sum := 0
	for nbr != 0 {
		digit := nbr % 10
		sum += digit * digit
		nbr /= 10
	}
	squareop <- sum
}

func calcCubes(nbr int, cubeop chan int) {
	sum := 0
	for nbr != 0 {
		digit := nbr % 10
		sum += digit * digit * digit
		nbr /= 10
	}
	cubeop <- sum
}

func main() {
	nbr := 589
	squarech := make(chan int)
	cubech := make(chan int)

	go calcSquares(nbr, squarech)
	go calcCubes(nbr, cubech)

	squares, cubes := <-squarech, <-cubech

	fmt.Println("Final output: ", squares+cubes)
}
