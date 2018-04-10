package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid, error := uuid.NewV4()
	if error != nil {
		log.Println(error)
	}
	return scope.SetColumn("Id", uuid.String())
}
