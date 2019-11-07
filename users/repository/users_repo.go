package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	pb "github.com/muhammadhidayah/users-service/proto/users"
	"github.com/muhammadhidayah/users-service/users"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) users.Repository {
	return &userRepo{db}
}

// This function will insert user to database using gorm
func (repo *userRepo) CreateUser(user *pb.User) error {
	err := repo.db.Create(user).Error
	return err
}

// this function will update data user using gorm
// return err
func (repo *userRepo) UpdateUser(user *pb.User) error {
	// get ID as params in condition
	ID := user.Id
	fmt.Println(ID)

	// query/syntax to update
	err := repo.db.Model(pb.User{}).Where("id = ?", ID).Updates(user).Error
	return err
}

// this function will delete user using gorm,
// in fact this function will not really delete the record from the database
// just update the column to is_deleted to 1
// return error
func (repo *userRepo) DeleteUser(user *pb.User) error {
	// get ID as params in condition
	ID := user.Id

	// query/syntax to delete using gorm
	err := repo.db.Model(pb.User{}).Where("id = ?", ID).Update("is_deleted", 1).Error
	return err
}

// this function will find user by person_id
// return User and Error
func (repo *userRepo) GetUserByPersonID(user *pb.User) (*pb.User, error) {
	// declare variable with type User to contain output of query
	var users pb.User

	// get personID as params in condition gorm sql
	personID := user.PersonId

	err := repo.db.Where("person_id = ?", personID).Find(&users).Error
	user2 := &users
	return user2, err
}

// this function will find user by person_id and password
func (repo *userRepo) GetUserByPersonIDAndPassword(user *pb.User) (*pb.User, error) {
	// declare variable with type User to contain output of query
	var users *pb.User

	// get personID as paramas inn condition gorm sql
	personID := user.PersonId
	// get person_password as paramas inn condition gorm sql
	personPassword := user.PersonPassword

	// query/syntax to find user by person_id and password,
	// variable users will contain output of this query
	// and return error
	err := repo.db.Where("person_id = ? AND person_password = ?", personID, personPassword).Find(users).Error

	return users, err
}
