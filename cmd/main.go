package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/amitpanthi/joblin/internal/queue"
	"github.com/amitpanthi/joblin/internal/router"
	"github.com/amitpanthi/joblin/internal/worker"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	jq := queue.NewJobQueue(5)
	w := worker.SpawnWorker(1, jq.Jobs)

	go w.Start(ctx)

	httpServer := http.Server{
		Addr:    ":8080",
		Handler: router.NewRouter(jq),
	}

	go func() {
		httpServer.ListenAndServe()
	}()

	go func() {
		<-sigs
		fmt.Println("Closing the program.")
		cancel()
		httpServer.Close()
	}()

	<-ctx.Done()
}
