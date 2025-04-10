package worker

import (
	"fmt"
	"sync"

	"github.com/amitpanthi/joblin/job"
)

type Worker struct {
	ID      int
	JobChan chan job.Job
}

func SpawnWorker(id int, jobChan chan job.Job) *Worker {
	return &Worker{
		ID:      id,
		JobChan: jobChan,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range w.JobChan {
		handle(job)
	}
}

func handle(job job.Job) {
	fmt.Printf("Job %d handled, job type - %s.\n", job.ID, job.Type)
}
