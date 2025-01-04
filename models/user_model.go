package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID        int
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u *UserModel) BeforeCreate(scope *gorm.DB) error {
	fmt.Println("BeforeCreate钩子函数")
	return nil
}

func (u *UserModel) BeforeUpdate(scope *gorm.DB) error {
	fmt.Println("BeforeUpdate")
	return nil
}

func (u *UserModel) BeforeDelete(scope *gorm.DB) error {
	fmt.Println("BeforeDelete")
	return nil
}