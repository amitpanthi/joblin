package api

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/amitpanthi/joblin/internal/job"
	"github.com/amitpanthi/joblin/internal/queue"
	"github.com/gin-gonic/gin"
)

var counter atomic.Int64

type handler struct {
	JobQueue *queue.JobQueue
}

type NewJob struct {
	Type    string         `json:"type"`
	Payload map[string]any `json:"payload"`
}

func NewHandler(jq *queue.JobQueue) *handler {
	return &handler{
		JobQueue: jq,
	}
}

func (h *handler) CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (h *handler) PostNewJob(c *gin.Context) {
	var newJob NewJob
	if err := c.ShouldBindJSON(&newJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse the input json: " + err.Error(),
		})
	}

	job := convertNewJobToJob(&newJob)
	h.JobQueue.PushJob(*job)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully pushed Job to queue with ID: %d", job.ID),
	})
}

func convertNewJobToJob(newJob *NewJob) *job.Job {
	return &job.Job{
		ID:      int(nextID()),
		Type:    newJob.Type,
		Payload: newJob.Payload,
	}
}

func nextID() int64 {
	return counter.Add(1)
}
