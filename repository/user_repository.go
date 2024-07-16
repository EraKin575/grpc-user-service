package repository

import (
    "errors"
    "grpc-user-service/grpc-user-service/proto"
)

type UserRepository interface {
    GetUserByID(id int32) (*proto.User, error)
    ListUsersByID(ids []int32) ([]*proto.User, error)
    
    SearchUsers(city string, phone int64, married bool) ([]*proto.User, error)
}

type InMemoryUserRepository struct {
    users []proto.User
}

func NewInMemoryUserRepository() UserRepository {
    return &InMemoryUserRepository{
        users: []proto.User{
            {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
            {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.1, Married: false},
            // Add more users as needed
        },
    }
}

func (repo *InMemoryUserRepository) GetUserByID(id int32) (*proto.User, error) {
    for _, user := range repo.users {
        if user.Id == id {
            return &user, nil
        }
    }
    return nil, errors.New("user not found")
}

func (repo *InMemoryUserRepository) ListUsersByID(ids []int32) ([]*proto.User, error) {
    var userList []*proto.User
    for _, id := range ids {
        for _, user := range repo.users {
            if user.Id == id {
                userList = append(userList, &user)
            }
        }
    }
    return userList, nil
}

func (repo *InMemoryUserRepository) SearchUsers(city string, phone int64, married bool) ([]*proto.User, error) {
    var userList []*proto.User
    for _, user := range repo.users {
        if (city == "" || user.City == city) &&
           (phone == 0 || user.Phone == phone) &&
           (married == user.Married) {
            userList = append(userList, &user)
        }
    }
    return userList, nil
}
