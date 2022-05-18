package rest_proxy

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
	"github.com/jtomasevic/gRPCExample/pkg/services/tasks/config"
	"google.golang.org/grpc"
)

var (
	// command-line options:
	// gRPC server endpoint
	tasksGrpcServerEndpoint = flag.String("tasks-grpc-server-endpoint", "localhost:"+config.GrpcPort, "gRPC server endpoint")
)

func RunProxy(ctx context.Context) {

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := service.RegisterTaskServiceHandlerFromEndpoint(ctx, mux, *tasksGrpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting REST Proxy server")
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	http.ListenAndServe(":"+config.ProxyPort, mux)
}
