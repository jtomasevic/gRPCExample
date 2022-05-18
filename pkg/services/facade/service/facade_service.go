package service

import (
	"context"

	tasks "github.com/jtomasevic/gRPCExample/pkg/services/facade/clients/tasks"
	users "github.com/jtomasevic/gRPCExample/pkg/services/facade/clients/users"
	taskService "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	userService "github.com/jtomasevic/gRPCExample/pkg/services/users/service"
	"github.com/jtomasevic/gRPCExample/pkg/workflows"
)

type facadeServiceImpl struct {
	userClient userService.UserServiceClient
	taskClient taskService.TaskServiceClient
	workFlows  workflows.WorkFlows
}

func (service *facadeServiceImpl) Signup(context context.Context, request *SignupRequset) (*SignupResponse, error) {
	result, err := service.userClient.Signup(context, &userService.SignupRequset{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return &SignupResponse{
			IsOk: false,
		}, err
	}
	return &SignupResponse{
		IsOk: true,
		Id:   result.Id,
	}, nil
}
func (service *facadeServiceImpl) AllUsers(context context.Context, request *GetAllUsersRequest) (*GetAlllUsersResponse, error) {
	result, _ := service.userClient.All(context, &userService.GetAllUsersRequest{})
	var users = make([]*User, len(result.Users))
	for i, u := range result.Users {
		users[i] = &User{
			Email:         u.Email,
			Id:            u.Id,
			NumberOfTasks: u.NumberOfTasks,
		}
	}
	return &GetAlllUsersResponse{
		Users: users,
	}, nil
}

func (service *facadeServiceImpl) CreateTask(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	workFlowRequest := &taskService.CreateRequest{
		Task: &taskService.Task{
			Title:  request.Task.Title,
			UserId: request.Task.UserId,
		},
	}
	response, _ := service.workFlows.CreateTaskWorkFlow.Run(ctx, workFlowRequest)
	return &CreateResponse{
		Id: response.Id,
	}, nil
}

// Read task
func (service *facadeServiceImpl) GetTaskById(ctx context.Context, request *ReadRequest) (*ReadResponse, error) {
	return &ReadResponse{}, nil
}

// Update task
func (service *facadeServiceImpl) UpdateTask(ctx context.Context, request *UpdateRequest) (*UpdateResponse, error) {
	return &UpdateResponse{}, nil
}

// Delete task
func (service *facadeServiceImpl) DeleteTask(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	return &DeleteResponse{}, nil
}

// Read all tasks
func (service *facadeServiceImpl) ReadAllTasks(ctx context.Context, request *ReadAllRequest) (*ReadAllResponse, error) {
	response, _ := service.taskClient.ReadAll(ctx, &taskService.ReadAllRequest{})
	var tasks = make([]*Task, len(response.Task))
	for i, t := range response.Task {
		tasks[i] = &Task{
			Id:          t.Id,
			Title:       t.Title,
			UserId:      t.UserId,
			Description: t.Description,
		}
	}
	return &ReadAllResponse{
		Task: tasks,
	}, nil
}

func (service *facadeServiceImpl) mustEmbedUnimplementedFacadeServiceServer() {

}

func NewFacadeService(ctx context.Context) FacadeServiceServer {
	userClient := users.NewClient()
	taskClient := tasks.NewClient()
	workFlows := workflows.NewWorkFlows(userClient, taskClient)
	return &facadeServiceImpl{
		userClient: userClient,
		taskClient: taskClient,
		workFlows:  *workFlows,
	}
}
