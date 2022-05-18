package workflows

import "github.com/jtomasevic/gRPCExample/pkg/workflows/create_task"
import 	tasks "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
import users "github.com/jtomasevic/gRPCExample/pkg/services/users/service"
import temClient "github.com/jtomasevic/gRPCExample/pkg/workflows/temporal"


type WorkFlows struct {
	CreateTaskWorkFlow create_task.CreateTaskWorkFlow
}	

func NewWorkFlows(
	userClient users.UserServiceClient,
	tasksClient tasks.TaskServiceClient,
) *WorkFlows {
	workFlowClient := temClient.GetClient()
	return &WorkFlows{
		CreateTaskWorkFlow: create_task.NewCreateTaskWorkFlow(
			userClient,
			tasksClient,
			workFlowClient,
		),
	}
}