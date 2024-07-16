package service_test

import (
    "context"
    "grpc-user-service/grpc-user-service/proto"
    "grpc-user-service/repository"
    "grpc-user-service/service"
    "testing"
)

func TestGetUser(t *testing.T) {
    repo := repository.NewInMemoryUserRepository()
    userService := service.NewUserService(repo)

    t.Run("User exists", func(t *testing.T) {
        req := &proto.GetUserRequest{Id: 1}
        res, err := userService.GetUser(context.Background(), req)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }
        if res.User.Id != req.Id {
            t.Errorf("expected user ID %v, got %v", req.Id, res.User.Id)
        }
    })

    t.Run("User does not exist", func(t *testing.T) {
        req := &proto.GetUserRequest{Id: 99}
        _, err := userService.GetUser(context.Background(), req)
        if err == nil {
            t.Fatalf("expected an error, got nil")
        }
    })
}

func TestListUsers(t *testing.T) {
    repo := repository.NewInMemoryUserRepository()
    userService := service.NewUserService(repo)

    t.Run("Users exist", func(t *testing.T) {
        req := &proto.ListUsersRequest{Ids: []int32{1, 2}}
        res, err := userService.ListUsers(context.Background(), req)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }
        if len(res.Users) != 2 {
            t.Errorf("expected 2 users, got %v", len(res.Users))
        }
    })

    t.Run("Some users do not exist", func(t *testing.T) {
        req := &proto.ListUsersRequest{Ids: []int32{1, 99}}
        res, err := userService.ListUsers(context.Background(), req)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }
        if len(res.Users) != 1 {
            t.Errorf("expected 1 user, got %v", len(res.Users))
        }
    })
}

func TestSearchUsers(t *testing.T) {
    repo := repository.NewInMemoryUserRepository()
    userService := service.NewUserService(repo)

    t.Run("Search by city", func(t *testing.T) {
        req := &proto.SearchUsersRequest{City: "LA"}
        res, err := userService.SearchUsers(context.Background(), req)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }
        if len(res.Users) == 0 {
            t.Errorf("expected users in LA, got none")
        }
    })

    t.Run("Search by phone", func(t *testing.T) {
        req := &proto.SearchUsersRequest{Phone: 1234567890}
        res, err := userService.SearchUsers(context.Background(), req)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }
        if len(res.Users) == 0 {
            t.Errorf("expected users with phone number 1234567890, got none")
        }
    })

    t.Run("Search by marital status", func(t *testing.T) {
        req := &proto.SearchUsersRequest{Married: true}
        res, err := userService.SearchUsers(context.Background(), req)
        if err != nil {
            t.Fatalf("expected no error, got %v", err)
        }
        if len(res.Users) == 0 {
            t.Errorf("expected married users, got none")
        }
    })
}
