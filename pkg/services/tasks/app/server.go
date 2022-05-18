package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jtomasevic/gRPCExample/pkg/services/tasks/config"
	"github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo"
	ds "github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo/data_source"
	"github.com/jtomasevic/gRPCExample/pkg/services/tasks/repo/migrations"
	"github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	"google.golang.org/grpc"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	migrations.TasksDbAutoMigrate()

	dataSource := ds.NewTasksDataSource()
	taskRepo := repo.NewTaskRepository(dataSource)
	taskService := service.NewTaskService(taskRepo)

	return runServer(ctx, taskService, config.GrpcPort)
}

// RunServer runs gRPC service to publish ToDo service
func runServer(ctx context.Context, api service.TaskServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	service.RegisterTaskServiceServer(server, api)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("starting Tasks server at port: %s", port)
	return server.Serve(listen)
}

func GetServer() (*grpc.Server, net.Listener) {

	migrations.TasksDbAutoMigrate()

	dataSource := ds.NewTasksDataSource()
	taskRepo := repo.NewTaskRepository(dataSource)
	api := service.NewTaskService(taskRepo)

	// ctx := context.Background()
	listener, err := net.Listen("tcp", ":"+config.GrpcPort)
	if err != nil {
		return nil, nil
	}

	// register service
	server := grpc.NewServer()
	service.RegisterTaskServiceServer(server, api)
	return server, listener
}
