package create_task

import (
	"log"
	"time"

	taskService "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	userService "github.com/jtomasevic/gRPCExample/pkg/services/users/service"
	"go.temporal.io/sdk/temporal"
	wf "go.temporal.io/sdk/workflow"
)

func ExecuteCreateTaskWorkFlow(ctx wf.Context, request *taskService.CreateRequest) (*taskService.CreateResponse, error) {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    500,
	}
	options := wf.ActivityOptions{
		// Timeout options specify when to automatically timeout Actvitivy functions.
		StartToCloseTimeout: time.Minute,
		// Optionally provide a customized RetryPolicy.
		// Temporal retries failures by default, this is just an example.
		RetryPolicy: retrypolicy,
	}
	ctx = wf.WithActivityOptions(ctx, options)
	var result taskService.CreateResponse
	err := wf.ExecuteActivity(ctx, CreateTaskAction, request).Get(ctx, &result)
	if err != nil {
		return nil, err
	}
	log.Printf("---- result from task action: %s", &result)
	err = wf.ExecuteActivity(ctx, AddTaskAction, &userService.AddTaskRequest{
		UserId: request.Task.UserId,
		TaskId: request.Task.Id,
	}).Get(ctx, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("---- RETURN from task action: %s", &result)
	return &result, nil
}
