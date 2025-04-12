package router

import (
	"github.com/amitpanthi/joblin/internal/api"
	"github.com/amitpanthi/joblin/internal/queue"
	"github.com/gin-gonic/gin"
)

func NewRouter(jq *queue.JobQueue) *gin.Engine {
	h := api.NewHandler(jq)
	r := gin.Default()

	r.GET("/internal/health", h.CheckHealth)

	{
		apiGroup := r.Group("/api/v1")
		apiGroup.POST("/jobs", h.PostNewJob)
	}

	return r
}
