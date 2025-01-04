package main

import (
	"fmt"
	"gorm_studyDemo/global"
	"time"
)

type UserModel struct {
	ID        int
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18"`
	CreatedAt time.Time
}

func migrate() {
	err := global.DB.AutoMigrate(&UserModel{})
	if err != nil {
		fmt.Println("表结构迁移失败")
		return
	}
	fmt.Println("表结构迁移成功")
}

func main() {
	global.Connect()

	migrate()

}
