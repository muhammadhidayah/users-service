package grpc

import (
	"context"

	pb "github.com/muhammadhidayah/users-service/proto/users"
	"github.com/muhammadhidayah/users-service/users"
)

// this struct will implement all method interface of UsersServiceHandler
type UserHandlerGRPC struct {
	UserUC users.Usecase
}

// implement CreateUser method of UsersServiceHandler interface
func (handler *UserHandlerGRPC) CreateUser(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := handler.UserUC.CreateUser(req); err != nil {
		res.Created = false
		return err
	}

	res.Created = true
	return nil
}

// implement UpdateUser method of UsersServiceHandler interface
func (handler *UserHandlerGRPC) UpdateUser(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := handler.UserUC.UpdateUser(req); err != nil {
		res.Updated = false
		return err
	}

	res.Updated = true
	res.User = req
	return nil
}

// implement DeleteUser method of UsersServiceHandler interface
func (handler *UserHandlerGRPC) DeleteUser(ctx context.Context, req *pb.User, res *pb.Response) error {
	if err := handler.UserUC.DeleteUser(req); err != nil {
		res.Deleted = false
		return err
	}

	res.Deleted = true
	return nil
}

// implement GetUserByPersonID method of UsersServiceHandler interface
func (handler *UserHandlerGRPC) GetUserByPersonID(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := handler.UserUC.GetUserByPersonID(req)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}

// implement GetUserByPersonIDAndPassword method of UsersServiceHandler interface
func (handler *UserHandlerGRPC) GetUserByPersonIDAndPassword(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := handler.UserUC.GetUserByPersonIDAndPassword(req)
	if err != nil {
		return err
	}

	res.User = user
	return nil
}
