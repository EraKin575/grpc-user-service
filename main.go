package main

import (
	"fmt"
	"grpc-user-service/grpc-user-service/proto"
	"grpc-user-service/repository"
	"grpc-user-service/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    userRepository := repository.NewInMemoryUserRepository()
    userService := service.NewUserService(userRepository)

    proto.RegisterUserServiceServer(grpcServer, userService)

    fmt.Println("Server is running on port :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
