package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm_studyDemo/models"
	"log"
)

var DB *gorm.DB

func Migrate() {
	err := DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatalf("数据库迁移失败 %s", err)
	}
	fmt.Println("数据库迁移成功")
}

func Connect() {
	dsn := "root:8888.216@tcp(127.0.0.1:3306)/gorm_new_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
