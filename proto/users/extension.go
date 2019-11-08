package user

import (
	"github.com/jinzhu/gorm"
	uuidgo "github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuidgo.NewV4()
	return scope.SetColumn("Id", uuid.String())
}
