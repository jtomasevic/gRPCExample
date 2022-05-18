package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	facade "github.com/jtomasevic/gRPCExample/pkg/services/facade/app"
	facadeProxy "github.com/jtomasevic/gRPCExample/pkg/services/facade/rest_proxy/proxy"
	task "github.com/jtomasevic/gRPCExample/pkg/services/tasks/app"
	tasksProxy "github.com/jtomasevic/gRPCExample/pkg/services/tasks/rest_proxy"
	user "github.com/jtomasevic/gRPCExample/pkg/services/users/app"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	go registerGrpcServer(ctx, task.GetServer, "Task")
	go registerGrpcServer(ctx, user.GetServer, "User")
	go registerGrpcServer(ctx, facade.GetServer, "Facade")
	go facadeProxy.RunProxy(ctx)
	go tasksProxy.RunProxy(ctx)
	// fmt.Println("all servers are started")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}

func registerGrpcServer(ctx context.Context, getServer func() (*grpc.Server, net.Listener), serverName string) {
	server, listener := getServer()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Printf("Shutting down %s server...\n", serverName)
			server.GracefulStop()
			fmt.Printf("%s server is closed\n", serverName)
			<-ctx.Done()
		}
	}()
	log.Printf("Starting %s server\n", serverName)
	if err := server.Serve(listener); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	} else {
		log.Printf("%s server started\n", serverName)
	}
}
