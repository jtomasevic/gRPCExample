package service

import (
	"context"
	"errors"

	model "github.com/jtomasevic/gRPCExample/pkg/services/users/model"
	repo "github.com/jtomasevic/gRPCExample/pkg/services/users/repo"
)

type userServiceImpl struct {
	userRepo repo.UsersRepository
}

func (service *userServiceImpl) Signup(context context.Context, request *SignupRequset) (*SignupResponse, error) {
	var user model.User = model.User{
		Email:    request.Email,
		Password: request.Password,
	}
	user, _ = service.userRepo.Save(user)
	return &SignupResponse{
		IsOk: true,
		Id:   user.Id,
	}, nil
}

func (service *userServiceImpl) mustEmbedUnimplementedUserServiceServer() {

}

func (service *userServiceImpl) All(context.Context, *GetAllUsersRequest) (*ReadAllUsersResponse, error) {
	users, _ := service.userRepo.All()
	var usersResponse = make([]*User, len(users))
	for i, user := range users {
		usersResponse[i] = &User{
			Email:         user.Email,
			Id:            user.Id,
			NumberOfTasks: int32(user.NumberOfTasks),
		}
	}
	return &ReadAllUsersResponse{
		Users: usersResponse,
	}, nil
}

func (service *userServiceImpl) AddTask(ctx context.Context, request *AddTaskRequest) (*AddTaskResponse, error) {
	user, err := service.userRepo.GetByID(request.UserId)
	if err != nil {
		return nil, errors.New("error on GetByID")
	}
	user.NumberOfTasks += 1
	service.userRepo.Save(user)
	return &AddTaskResponse{
		UserName:        user.Email,
		LastTaskAddedId: request.TaskId,
		HasTasks:        int32(user.NumberOfTasks),
	}, nil
}

func NewUserService(userRepo repo.UsersRepository) UserServiceServer {
	return &userServiceImpl{
		userRepo,
	}
}
