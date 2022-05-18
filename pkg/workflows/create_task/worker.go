package create_task

import "go.temporal.io/sdk/client"
import temClient "github.com/jtomasevic/gRPCExample/pkg/workflows/temporal"

func RegisterWorker(client client.Client) {
	temClient.RegisterWorker(client, TaskQueues, ExecuteCreateTaskWorkFlow, CreateTaskAction, AddTaskAction)
}