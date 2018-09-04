package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended \n", i)
	wg.Done()
}

// Worker pool
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)

	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}

	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}

	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}

	done <- true
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()
	fmt.Println("All go routines have finished")

	// Worker pool
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 100
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}
