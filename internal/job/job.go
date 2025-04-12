package job

type Job struct {
	ID      int
	Type    string
	Payload map[string]any
}

func NewJob(id int, jobType string) *Job {
	return &Job{
		ID:      id,
		Type:    jobType,
		Payload: make(map[string]any),
	}
}
