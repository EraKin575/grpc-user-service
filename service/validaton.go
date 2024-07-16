package service

import (
    "errors"
    "grpc-user-service/grpc-user-service/proto"
)

func validateGetUserRequest(req *proto.GetUserRequest) error {
    if req.Id <= 0 {
        return errors.New("invalid user ID")
    }
    return nil
}

func validateListUsersRequest(req *proto.ListUsersRequest) error {
    if len(req.Ids) == 0 {
        return errors.New("no user IDs provided")
    }
    for _, id := range req.Ids {
        if id <= 0 {
            return errors.New("invalid user ID")
        }
    }
    return nil
}

func validateSearchUsersRequest(req *proto.SearchUsersRequest) error {
    if req.City == "" && req.Phone == 0 && !req.Married {
        return errors.New("at least one search criteria must be provided")
    }
    return nil
}
