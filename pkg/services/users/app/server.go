package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jtomasevic/gRPCExample/pkg/services/users/repo"
	migrations "github.com/jtomasevic/gRPCExample/pkg/services/users/repo/migrations"
	ds "github.com/jtomasevic/gRPCExample/pkg/services/users/repo/data_source"

	"github.com/jtomasevic/gRPCExample/pkg/services/users/service"
	"google.golang.org/grpc"
)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	migrations.UsersDbAutoMigrate()

	dataSource := ds.NewUsersDataSource()
	repo := repo.NewUserRepository(dataSource)
	userService := service.NewUserService(repo)

	return runServer(ctx, userService, "50060")
}

// RunServer runs gRPC service to publish ToDo service
func runServer(ctx context.Context, api service.UserServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	service.RegisterUserServiceServer(server, api)

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
	log.Printf("Starting Tasks server at tcp port %s\n", port)

	return server.Serve(listen)
}


// RunServer runs gRPC service to publish ToDo service
func GetServer() (*grpc.Server, net.Listener) {
	migrations.UsersDbAutoMigrate()

	dataSource := ds.NewUsersDataSource()
	repo := repo.NewUserRepository(dataSource)
	api := service.NewUserService(repo)
	// ctx := context.Background()
	port := "50060"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, nil
	}

	// register service
	server := grpc.NewServer()
	service.RegisterUserServiceServer(server, api)
	return server, listener
}