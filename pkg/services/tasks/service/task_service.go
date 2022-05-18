package service

import (
	"context"
	"fmt"

	model "github.com/jtomasevic/gRPCExample/pkg/services/tasks/model"
	repo "github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo"
)

type taskServiceImpl struct {
	taskRepo repo.TasksRepository
}

func (service *taskServiceImpl) Create(ctx context.Context, request *CreateRequest) (*CreateResponse, error) {
	var task model.Task = model.Task{
		Title: request.Task.Title,
	}
	task, _ = service.taskRepo.Save(task)
	return &CreateResponse{
		Id: task.Id,
	}, nil
}
func (service *taskServiceImpl) GetById(ctx context.Context, request *ReadRequest) (*ReadResponse, error) {
	task, _ := service.taskRepo.GetByID(request.Id)
	return &ReadResponse{
		Task: mapToViewModel(task),
	}, nil
}

func (service *taskServiceImpl) Update(ctx context.Context, request *UpdateRequest) (*UpdateResponse, error) {
	task, _ := service.taskRepo.GetByID(request.Task.Id)
	task.Title = request.Task.Title
	task.Description = request.Task.Description
	service.taskRepo.Save(task)
	return &UpdateResponse{
		Updated: 1,
	}, nil
}

func (service *taskServiceImpl) Delete(ctx context.Context, request *DeleteRequest) (*DeleteResponse, error) {
	service.taskRepo.Delete(request.Id)
	return &DeleteResponse{
		Deleted: 1,
	}, nil
}

func (service *taskServiceImpl) ReadAll(context.Context, *ReadAllRequest) (*ReadAllResponse, error) {
	tasks, _ := service.taskRepo.All()
	var viewTasks = make([]*Task, len(tasks))
	for i, task := range tasks {
		viewTasks[i] = mapToViewModel(task)
	}
	return &ReadAllResponse{
		Task: viewTasks,
	}, nil
}

func (service *taskServiceImpl) mustEmbedUnimplementedTaskServiceServer() {
	fmt.Println("!!!! mustEmbedUnimplementedTaskServiceServer")
}

func NewTaskService(taskRepo repo.TasksRepository) TaskServiceServer {
	return &taskServiceImpl{
		taskRepo,
	}
}

func mapToViewModel(task model.Task) *Task {
	return &Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		UserId:      task.UserId,
	}
}
