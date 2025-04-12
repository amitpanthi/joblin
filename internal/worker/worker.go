package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/amitpanthi/joblin/internal/job"
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

func (w *Worker) Start(ctx context.Context) {
	for {
		select {
		case job := <-w.JobChan:
			handle(job)
		case <-ctx.Done():
			return
		}
	}
}

func handle(job job.Job) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Job %d handled, job type - %s.\n", job.ID, job.Type)
}
