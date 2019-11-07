package usecase

import (
	"errors"
	"fmt"

	pb "github.com/muhammadhidayah/users-service/proto/users"
	"github.com/muhammadhidayah/users-service/users"
	"golang.org/x/crypto/bcrypt"
)

type usersUsecase struct {
	repo users.Repository
}

func NewUsersUsecase(repo users.Repository) users.Usecase {
	return &usersUsecase{repo}
}

// this function to handling all process bussiness related of create user
func (usecase *usersUsecase) CreateUser(user *pb.User) error {

	// Generates a hash password using bcrypt
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.PersonPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New(fmt.Sprintf("error hashing password : %v", err))
	}

	// after we hashed the password, then we store the hashed password to user.Password
	// and dont forget to casting type data to string
	user.PersonPassword = string(hashedPass)
	if err := usecase.repo.CreateUser(user); err != nil {
		return errors.New(fmt.Sprintf("error creating user: %v", err))
	}

	return nil
}

// this function to handling all process bussiness related of update user
// return error
func (usecase *usersUsecase) UpdateUser(user *pb.User) error {
	// personPassword will hashed, because the user updated their password
	if user.PersonPassword != "" {
		// generate new hashing for their password
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.PersonPassword), bcrypt.DefaultCost)
		if err != nil {
			return errors.New(fmt.Sprintf("error hashing password : %v", err))
		}

		user.PersonPassword = string(hashedPass)
	}

	if err := usecase.repo.UpdateUser(user); err != nil {
		return errors.New(fmt.Sprintf("error updating user: %v", err))
	}

	return nil
}

// this function to delete user
func (usecase *usersUsecase) DeleteUser(user *pb.User) error {
	if err := usecase.repo.DeleteUser(user); err != nil {
		return errors.New(fmt.Sprintf("error to delete user: %v", err))
	}

	return nil
}

func (usecase *usersUsecase) GetUserByPersonID(user *pb.User) (*pb.User, error) {
	res, err := usecase.repo.GetUserByPersonID(user)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (usecase *usersUsecase) GetUserByPersonIDAndPassword(user *pb.User) (*pb.User, error) {
	// we will get user by personID first, and then we comparing password input and password to DB
	res, err := usecase.GetUserByPersonID(user)
	if err != nil {
		return nil, err
	}

	// we comparing betweer passwordDB and password user input
	if err := bcrypt.CompareHashAndPassword([]byte(res.PersonPassword), []byte(user.PersonPassword)); err != nil {
		return nil, err
	}

	return res, nil
}
