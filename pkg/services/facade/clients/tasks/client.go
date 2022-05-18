package client

import (
	"log"

	"google.golang.org/grpc"

	srv "github.com/jtomasevic/gRPCExample/pkg/services/tasks/service"
)

func NewClient() srv.TaskServiceClient {
	adr := "localhost:50055"
	address := &adr
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()

	return srv.NewTaskServiceClient(conn)
}
