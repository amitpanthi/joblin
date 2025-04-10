package main

import (
	"sync"

	"github.com/amitpanthi/joblin/job"
	"github.com/amitpanthi/joblin/queue"
	"github.com/amitpanthi/joblin/worker"
)

func main() {
	var wg sync.WaitGroup

	jq := queue.NewJobQueue(5)
	w := worker.SpawnWorker(1, jq.Jobs)

	wg.Add(1)
	go w.Start(&wg)
	jq.PushJob(*job.NewJob(1, "Test job 1"))
	jq.PushJob(*job.NewJob(2, "Test job 2"))
	jq.PushJob(*job.NewJob(3, "Test job 3"))

	close(jq.Jobs)
	wg.Wait()
}
