package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "grpc-user-service/grpc-user-service/proto"
    "grpc-user-service/repository"
    "grpc-user-service/service"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    userService := service.NewUserService(repository.NewInMemoryUserRepository())
    proto.RegisterUserServiceServer(grpcServer, userService)

    // Register reflection service on gRPC server
    reflection.Register(grpcServer)

    log.Println("Server is running on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
