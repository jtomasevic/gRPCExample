package client

import (
	"log"

	"google.golang.org/grpc"

	srv "github.com/jtomasevic/gRPCExample/pkg/services/users/service"
)

func NewClient() srv.UserServiceClient {
	adr := "localhost:50060"
	address := &adr
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()

	return srv.NewUserServiceClient(conn)
}
