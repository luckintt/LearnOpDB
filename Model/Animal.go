package Model

import (
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jinzhu/gorm"
)

type Animal struct {
	AnimalID string `gorm:"primary_key" json:"animal_id"`
	MasterId int 	`json:"master_id"`
	Name     string	`json:"name"`
	Age      int	`json:"age"`
}

func (a *Animal) TableName() string {
	return "animals"
}

func (a *Animal) BeforeCreate(scope *gorm.Scope) error {
	UUID, _ := uuid.NewV4()
	scope.SetColumn("AnimalID", UUID.String())
	return nil
}
