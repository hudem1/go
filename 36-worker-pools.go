// Nothing really new (technically) in this file actually, more like a concept.
//
// The worker pool pattern is a design in which a fixed number of workers are given a stream of tasks to process in a queue.
// The tasks stay in the queue until a worker is free to pick it up and execute it.
package main

import (
	"fmt"
	"time"
)

// A worker receives work on the `jobs` channel and send the corresponding result on the `results` channel.
// We make workers sleep 1s per job to simulate an expensive task.
func workerCh36(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func worker_pools_main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Run several concurrent instances of `worker`.
	// Start up 3 workers, initially blocked (on trying to receive) because there are no jobs yet.
	for w := 1; w <= 3; w++ {
		go workerCh36(w, jobs, results)
	}

	// Send 5 jobs and then close the channel ton indicate that's all the work we have.
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// As a reminder, closing the channel allows to terminate the range loop in `worker`,
	// making goroutines exit the `worker` function after finishing all jobs
	// and not block themselves on trying to receive more jobs.
	close(jobs)

	// Wait (and optionally collect result) for all workers to finish their job.
	// We'll see an alternative way in next chapter: `WaitGroup`.
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
