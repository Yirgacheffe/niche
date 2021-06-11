package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id      int
	randNbr int
}

type Result struct {
	job         Job
	sumOfDigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	nbr := number
	for nbr != 0 {
		digit := nbr % 10
		sum += digit
		nbr /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func createWorkerPool(nbrOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < nbrOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randNbr)}
		results <- output
	}
	wg.Done()
}

func allocate(nbrOfJobs int) {
	for i := 0; i < nbrOfJobs; i++ {
		rdnNo := rand.Intn(999)
		job := Job{i, rdnNo}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random number %d, sum of digits %d\n", result.job.id, result.job.randNbr, result.sumOfDigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	nbrOfJobs := 100
	go allocate(nbrOfJobs)

	done := make(chan bool)
	go result(done)

	nbrOfWorkers := 20
	createWorkerPool(nbrOfWorkers)

	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)

	fmt.Println("Total time taken ", diff.Seconds(), "seconds")
}
