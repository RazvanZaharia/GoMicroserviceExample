package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid1 := uuid.NewV4()
	/*if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return nil
	}*/

	return scope.SetColumn("Id", uuid1.String())
}
