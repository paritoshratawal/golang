package main

import (
	"fmt"
	"sync"
)

type JOB struct {
	ID int
}

type Result struct {
	Job    JOB
	Output string
}

func worker(worker_id int, wg *sync.WaitGroup, jobs <-chan JOB, results chan<- Result) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", worker_id)
	for job := range jobs {
		results <- Result{
			Job:    job,
			Output: fmt.Sprintf("Worker %d processed job %d", worker_id, job.ID),
		}
	}
}

func main() {
	const (
		numJobs    = 10
		numWorkers = 3
	)
	wg := &sync.WaitGroup{}

	jobs := make(chan JOB, numJobs)
	results := make(chan Result, numJobs)

	for worker_id := 1; worker_id <= numWorkers; worker_id++ {
		wg.Add(1)
		go worker(worker_id, wg, jobs, results)
	}

	for job_id := 1; job_id <= numJobs; job_id++ {
		jobs <- JOB{ID: job_id}
	}

	close(jobs)
	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result.Output)
	}
}
