package queue

import "github.com/amitpanthi/joblin/job"

type JobQueue struct {
	Jobs chan job.Job
}

func NewJobQueue(bufferSize int) *JobQueue {
	return &JobQueue{
		make(chan job.Job, bufferSize),
	}
}

func (jq *JobQueue) PushJob(newJob job.Job) {
	jq.Jobs <- newJob
}
