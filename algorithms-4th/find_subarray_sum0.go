package main

import "fmt"

func findSubarraySum0(a []int) {

	n := len(a)
	sum := 0
	hm := make(map[int][]int)

	for i := 0; i < n; i++ {
		sum += a[i]

		// Handle the special case sum = 0
		if sum == 0 {
			printRange(0, i)
		} else {
			if v, ok := hm[sum]; ok {

				for _, vv := range v {
					start := vv + 1
					end := i
					printRange(start, end)
				}

				v = append(v, i)
				hm[sum] = v

			} else {
				var t []int
				t = append(t, i)
				hm[sum] = t
			}
		}
	}

}

func printRange(start, end int) {
	fmt.Printf("Subarray found from Index %d to %d\n", start, end)
}

func main() {
	a := []int{6, 3, -1, -3, 4, -2, 2, 4, 6, -12, -7}
	findSubarraySum0(a)
}
