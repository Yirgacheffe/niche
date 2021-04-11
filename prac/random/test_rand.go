package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano())

	rdn := rand.Intn(30)

	if rdn <= 10 {
		fmt.Println("kdjsf;skd")
	} else if rdn <= 20 {
		fmt.Println("kkkkkkkk")
	} else {
		fmt.Println("kdkdkdkdkdkdk")
	}

}
