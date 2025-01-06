package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID              int
	Name            string           `gorm:"size:16;not null;unique"`
	Age             int              `gorm:"default:18"`
	UserDetailModel *UserDetailModel `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

type UserDetailModel struct {
	ID        int
	UserID    int        `grom:"unique"`
	UserModel *UserModel `gorm:"foreignkey:UserID;constraint:OnDelete:CASCADE"`
	Email     string     `gorm:"size:32"`
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

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("查询钩子")
	return
}
