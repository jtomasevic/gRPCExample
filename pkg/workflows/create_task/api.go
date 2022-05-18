package create_task

import (
	"context"

	tasks "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	users "github.com/jtomasevic/gRPCExample/pkg/services/users/service"
	"go.temporal.io/sdk/client"
)

type CreateTaskWorkFlow interface {
	Run(ctx context.Context, request *tasks.CreateRequest) (*tasks.CreateResponse, error)
}

type createTaskWorkFlow struct {
	tasksClient    tasks.TaskServiceClient
	userClient     users.UserServiceClient
	workFlowClient client.Client
}

func NewCreateTaskWorkFlow(
	userClient users.UserServiceClient,
	tasksClient tasks.TaskServiceClient,
	workFlowClient client.Client,

) *createTaskWorkFlow {
	userClient = userClient
	taskClient = tasksClient
	RegisterWorker(workFlowClient)
	return &createTaskWorkFlow{
		tasksClient:    taskClient,
		userClient:     userClient,
		workFlowClient: workFlowClient,
	}
}
