package main

import (
	"hello-world-temporal/internal/helloworld"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, helloworld.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(helloworld.GreetingWorkflow)
	w.RegisterActivity(helloworld.ComposeGreeting)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
