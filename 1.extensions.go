package main

import "github.com/jinzhu/gorm"
import "github.com/satori/go.uuid"

func (model *User) BeforeCreate(scope *gorm.Scope) error {
	xuuid := uuid.NewV4()
	return scope.SetColumn("Id", xuuid.String())
}
