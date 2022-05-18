package temporal

import (
	"fmt"
	"log"

	// wf "github.com/jtomasevic/gRPCExample/pkg/work_flow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

var (
	temporalClient client.Client = nil
)

func GetClient() client.Client {
	if temporalClient == nil {
		fmt.Println("--- Creating new temporal client")
		t, err := client.NewClient(client.Options{})
		temporalClient = t
		if err != nil {
			log.Fatalln("unable to create Temporal client", err)
			fmt.Printf("unable to create Temporal client %s", err.Error())
		}
	}
	return temporalClient
}

func RegisterWorker(client client.Client, queueName string, workflow interface{}, methods ...interface{}) {
	go func() {
		w := worker.New(client, queueName, worker.Options{})
		w.RegisterWorkflow(workflow)
		for _, a := range methods {
			w.RegisterActivity(a)
		}
		// Start listening to the Task Queue
		err := w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("unable to start Worker", err)
		}
	}()
}
