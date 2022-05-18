package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/jtomasevic/gRPCExample/pkg/services/facade/service"
	"google.golang.org/grpc"

)

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	facadeService := service.NewFacadeService(ctx)

	return runServer(ctx, facadeService, "50010")
}

// RunServer runs gRPC service to publish ToDo service
func runServer(ctx context.Context, api service.FacadeServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	service.RegisterFacadeServiceServer(server, api)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down facade server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("starting Facade server at port: %s", port)

	return server.Serve(listen)
}

// RunServer runs gRPC service to publish ToDo service
func GetServer() (*grpc.Server, net.Listener) {
	ctx := context.Background()
	api := service.NewFacadeService(ctx)
	// ctx := context.Background()
	port := "50010"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, nil
	}

	// register service
	server := grpc.NewServer()
	service.RegisterFacadeServiceServer(server, api)
	return server, listener
}
