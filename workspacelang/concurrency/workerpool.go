package main

import (
	"fmt"
	"math/rand"
	"time"
)

type jobDetail struct {
	workerName string
	jobId      int
	result     int
}

func workerPool(numWorkers int, jobs <-chan jobDetail, results chan<- jobDetail) {
	for i := 0; i < numWorkers; i++ {
		go worker(fmt.Sprintf("Worker %d", i), jobs, results)
	}
}

func worker(workerName string, jobs <-chan jobDetail, results chan<- jobDetail) {
	for j := range jobs {
		j.workerName = workerName
		results <- process(j)
	}
}

func process(jobDetail jobDetail) jobDetail {
	fmt.Printf("Worker (%v) processing job (%v)\n", jobDetail.workerName, jobDetail.jobId)
	// Simulate some work
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	// return job * 2
	jobDetail.result = jobDetail.jobId * 2
	return jobDetail
}

func ExecuteWorkerPool() {
	numJobs := 20
	jobs := make(chan jobDetail, numJobs)
	results := make(chan jobDetail, numJobs)

	// Start the worker pool
	workerPool(5, jobs, results)

	// Send jobs
	for i := 0; i < numJobs; i++ {
		jobs <- jobDetail{jobId: i}
	}
	close(jobs)

	// Collect results
	for i := 0; i < numJobs; i++ {
		result := <-results
		fmt.Printf("Result(%d) executed by Worker (%s)\n", result.result, result.workerName)
	}
	fmt.Println("Done...")
}
