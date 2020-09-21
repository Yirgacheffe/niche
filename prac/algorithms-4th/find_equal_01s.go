package main

import "fmt"

type PairIdx struct {
	start int
	end   int
}

func replaceZero(a []int) []int {

	tmp := []int{}

	// Initial the temp array
	for _, v := range a {
		if v == 0 {
			tmp = append(tmp, -1)
		} else {
			tmp = append(tmp, v)
		}
	}

	return tmp

}

func bruteForce(a []int) {

	var p PairIdx
	n := len(a)

	maxlen := 0

	for i := 0; i < n; i++ {
		sum := a[i]

		for j := i + 1; j < n; j++ {
			sum += a[j]

			if sum == 0 {
				tmpLen := j - i + 1
				if maxlen <= tmpLen {
					maxlen = tmpLen
					p = PairIdx{start: i, end: j}
				}
			}
		}

	}

	fmt.Printf("Max lenth is: %d\n", maxlen)
	fmt.Printf("Index is: %d -> %d\n", p.start, p.end)

}

func main() {
	a := []int{1, 0, 1, 1, 1, 0, 0}

	tmp := replaceZero(a)
	bruteForce(tmp)
}
