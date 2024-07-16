package service

import (
    "context"
    "grpc-user-service/grpc-user-service/proto"
    "grpc-user-service/repository"
    "log"
)

type UserService struct {
    proto.UnimplementedUserServiceServer
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
    if err := validateGetUserRequest(req); err != nil {
        return nil, err
    }
    log.Printf("GetUser called with ID: %v", req.Id)
    user, err := s.repo.GetUserByID(req.Id)
    if err != nil {
        log.Printf("Error getting user: %v", err)
        return nil, err
    }
    return &proto.GetUserResponse{User: user}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *proto.ListUsersRequest) (*proto.ListUsersResponse, error) {
    if err := validateListUsersRequest(req); err != nil {
        return nil, err
    }
    log.Printf("ListUsers called with IDs: %v", req.Ids)
    users, err := s.repo.ListUsersByID(req.Ids)
    if err != nil {
        log.Printf("Error listing users: %v", err)
        return nil, err
    }
    return &proto.ListUsersResponse{Users: users}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, req *proto.SearchUsersRequest) (*proto.SearchUsersResponse, error) {
    if err := validateSearchUsersRequest(req); err != nil {
        return nil, err
    }
    log.Printf("SearchUsers called with criteria: City=%v, Phone=%v, Married=%v", req.City, req.Phone, req.Married)
    users, err := s.repo.SearchUsers(req.City, req.Phone, req.Married)
    if err != nil {
        log.Printf("Error searching users: %v", err)
        return nil, err
    }
    return &proto.SearchUsersResponse{Users: users}, nil
}
