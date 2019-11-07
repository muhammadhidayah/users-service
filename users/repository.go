package users

import (
	pb "github.com/muhammadhidayah/users-service/proto/users"
)

type Repository interface {
	CreateUser(*pb.User) error
	UpdateUser(*pb.User) error
	DeleteUser(*pb.User) error
	GetUserByPersonID(*pb.User) (*pb.User, error)
	GetUserByPersonIDAndPassword(*pb.User) (*pb.User, error)
}
