package create_task

import (
	"context"
	"log"

	taskService "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	tasks "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	userService "github.com/jtomasevic/gRPCExample/pkg/services/users/service"
	"go.temporal.io/sdk/client"
)

var (
	taskClient taskService.TaskServiceClient
	userClient userService.UserServiceClient
)

const (
	WorkFlowId = "Create-task-alpha"
	TaskQueues = "Create-task-workflow"
)

func (w *createTaskWorkFlow) Run(ctx context.Context, request *tasks.CreateRequest) (*tasks.CreateResponse, error) {
	taskClient = w.tasksClient
	userClient = w.userClient
	options := client.StartWorkflowOptions{
		ID:        WorkFlowId,
		TaskQueue: TaskQueues,
	}

	// c := temClient.GetClient()
	// temClient.RegisterWorker(c, options.TaskQueue, CreateTaskWorkFlow, CreateTaskAction, AddTaskAction)
	workFlowRequest := &taskService.CreateRequest{
		Task: &taskService.Task{
			Title:  request.Task.Title,
			UserId: request.Task.UserId,
		},
	}
	we, err := w.workFlowClient.ExecuteWorkflow(ctx, options, ExecuteCreateTaskWorkFlow, workFlowRequest)

	if err != nil {
		log.Println("error starting TestWorkFlow workflow", err)
	} else {
		log.Printf("Workflow %s started", we.GetID())
	}
	var response taskService.CreateResponse
	err = we.Get(ctx, &response)
	if err != nil {
		log.Println("----- error get wf result", err)
	} else {
		log.Printf("---- result from wf %s", response.Id)
	}

	return &tasks.CreateResponse{
		Id: response.Id,
	}, nil
}

func CreateTaskAction(ctx context.Context, request *taskService.CreateRequest) (*taskService.CreateResponse, error) {
	task, _ := taskClient.Create(ctx, request)
	return task, nil
}

func AddTaskAction(ctx context.Context, in *userService.AddTaskRequest) error {
	userClient.AddTask(ctx, in)
	return nil
}
