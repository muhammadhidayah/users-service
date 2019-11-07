package repository

import (
	"github.com/jinzhu/gorm"
	pb "github.com/muhammadhidayah/users-service/proto/users"
	"github.com/muhammadhidayah/users-service/users"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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
	// get personID as params in condition
	personID := user.PersonId

	// query/syntax to update
	err := repo.db.Where("person_id = ?", personID).Update(user).Error
	return err
}

// this function will delete user using gorm,
// in fact this function will not really delete the record from the database
// just update the column to is_deleted to 1
// return error
func (repo *userRepo) DeleteUser(user *pb.User) error {
	// get personID as param in condition
	personID := user.PersonId
	// query/syntax to delete using gorm
	err := repo.db.Where("person_id = ?", personID).Update("is_deleted", 1).Error
	return err
}

// this function will find user by person_id
// return User and Error
func (repo *userRepo) GetUserByPersonID(user *pb.User) (*pb.User, error) {
	// declare variable with type User to contain output of query
	var users *pb.User

	// get personID as params in condition gorm sql
	personID := user.PersonId

	err := repo.db.Where("person_id = ?", personID).Find(users).Error
	return users, err
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
