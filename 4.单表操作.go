package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm_studyDemo/global"
	"gorm_studyDemo/models"
	"time"
)

func insert() {
	/**
	一次插入
	*/
	//err := global.DB.Create(&models.UserModel{
	//	Name: "firefire",
	//	Age:  25,
	//}).Error
	//
	//if err != nil {
	//	fmt.Println(err)
	//}

	/**
	回填式创建
	*/
	//user := models.UserModel{
	//	Name: "myl216",
	//	Age:  25,
	//}
	//
	//err := global.DB.Create(&user).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(user.ID, user.CreatedAt, user.Name)

	/**
	添加多个
	*/
	var userList = []models.UserModel{
		{Name: "myl1263"},
		{Name: "myl1234567"},
	}

	err := global.DB.Create(&userList).Error
	if err != nil {
		fmt.Println(err)
		return
	}
}

func query() {
	//var userList = []models.UserModel{}
	//err := global.DB.Find(&userList, "age >= ?", 25).Error
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(userList)

	var user models.UserModel
	// 根据主键查询
	//global.DB.Take(&user, 4)

	// 根据主键排序查第一个
	//global.DB.Debug().First(&user, 4)

	//user.ID = 6
	//global.DB.Take(&user)

	//err = global.DB.Take(&user, 111).Error
	//if err == gorm.ErrRecordNotFound {
	//	fmt.Println("记录不存在")
	//	return
	//}

	err := global.DB.Limit(1).Find(&user, 111).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("记录不存在")
		return
	}

	// .Debug() 打印实际的sql
	fmt.Println(user.ID == 0)
}

func updateSave() {
	var user models.UserModel
	user.ID = 4
	user.Name = "fire332313216"
	user.CreatedAt = time.Now()
	/**
	Save,有主键记录就是更新，并且可以更新零值，否则就是创建
	*/
	global.DB.Save(&user)
	fmt.Println(user)
}

func update() {
	var user = models.UserModel{ID: 3}
	//global.DB.Model(&user).Update("age", 0)
	/**
	与Update()唯一的区别就是不会走钩子函数
	*/
	global.DB.Model(&user).UpdateColumn("age", 22)
	fmt.Println(user)
}

func updates() {
	var user = models.UserModel{ID: 3}
	/**
	不会更新零值
	*/
	//global.DB.Model(&user).Updates(models.UserModel{
	//	Name: "fg",
	//	Age:  0,
	//})

	//global.DB.Model(&user).Updates(map[string]interface{}{
	//	"age": 0,
	//})
	//
	//fmt.Println(user)

	global.DB.Model(&user).Update("age", gorm.Expr("age+?", 1))

	fmt.Println(user)
}

func delete() {
	//var user = models.UserModel{ID: 8}
	//global.DB.Delete(&user)
	//fmt.Println(user)

	//global.DB.Delete(&models.UserModel{}, 7)
	//global.DB.Delete(&models.UserModel{},"name = ?","myl")

	/**
	批量删除
	*/
	//global.DB.Delete(&models.UserModel{}, []int{4, 5})

	var user models.UserModel
	//global.DB.Take(&user, 4)
	global.DB.Unscoped().Take(&user, 4)

	global.DB.Unscoped().Delete(&user)

	fmt.Println(user)

}

func HighQuery() {
	//var user models.UserModel
	//global.DB.Debug().Where("age > ?", 18).Take(&user)
	//fmt.Println(user)

	//var user models.UserModel
	//global.DB.Debug().Where(models.UserModel{
	//	Name: "firefire",
	//	Age:  0,
	//}).Take(&user)
	//fmt.Println(user)

	//var user models.UserModel
	//global.DB.Debug().Where(map[string]any{
	//	"age": 2,
	//}).Take(&user)
	//fmt.Println(user)

	//query := global.DB.Where("age > 18 and name = ?", "firefire")
	//var user models.UserModel
	//global.DB.Debug().Where(query).Take(&user)
	//fmt.Println(user)

	//var userList []models.UserModel
	//global.DB.Debug().Or("age = 18").Or("name = ?", "firefire").Find(&userList)
	//fmt.Println(userList)

	//var userList []models.UserModel
	//global.DB.Debug().Not("age > 18").Find(&userList)
	//fmt.Println(userList)

	//var userList []models.UserModel
	//global.DB.Debug().Order("age desc").Find(&userList)
	////global.DB.Debug().Order("age asc").Order("id desc").Find(&userList)
	//fmt.Println(userList)

	//var nameList []string
	//global.DB.Model(models.UserModel{}).Select("name").Scan(&nameList)
	//fmt.Println(nameList)

	//var nameList []string
	//global.DB.Model(models.UserModel{}).Pluck("name", &nameList)
	//fmt.Println(nameList)

	//type User struct {
	//	Name string `gorm:"column:name"`
	//	Age  int    `gorm:"column:age"`
	//}
	//
	//var userList []User
	//global.DB.Model(models.UserModel{}).Scan(&userList)
	//fmt.Println(userList)

	//type UserCount struct {
	//	Age    int
	//	Counut int
	//}
	//
	//var list []UserCount
	//global.DB.Model(models.UserModel{}).Group("age").Select("age", "count(id) as count").Scan(&list)
	//fmt.Println(list)

	//var ageList []int
	//global.DB.Model(models.UserModel{}).Distinct("age").Pluck("age", &ageList)
	//fmt.Println(ageList)

	//var userList []models.UserModel

	//global.DB.Limit(2).Offset(0).Find(&userList)
	//fmt.Println(userList)
	//
	//global.DB.Limit(2).Offset(2).Find(&userList)
	//fmt.Println(userList)
	//
	//global.DB.Limit(2).Offset(4).Find(&userList)
	//fmt.Println(userList)

	//limit := 2
	//page := 1
	//
	//global.DB.Limit(limit).Offset((page - 1) * limit).Find(&userList)
	//fmt.Println(userList)

	//var userList []models.UserModel
	////global.DB.Scopes(Age18).Find(&userList)
	//global.DB.Scopes(NameIn("firefire", "fg")).Find(&userList)
	//fmt.Println(userList)

	type User struct {
		Name string
		Age  int
	}

	var users []User

	global.DB.Raw("select name,age from user_models").Scan(&users)
	fmt.Println(users)

	global.DB.Exec("update user_models set age=? where id=?", 22, 3)

}

func NameIn(nameList ...string) func(db *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Where("name in ?", nameList)
	}
}

func Age18(tx *gorm.DB) *gorm.DB {
	return tx.Where("age > 18")
}

//更新
//有很多方法，Save、Update、UpdateColumn、Updates
//
//不同的方法有不同的区别，如下：
//
//Save，有主键记录就是更新，并且可以更新零值，否则就是创建
//Update，可以更新零值，必须要有条件
//UpdateColumn，可以更新零值，不会走更新的Hook
//Updates，如果是结构体，则更新非零值，map可以更新零值

func main() {
	global.Connect()
	//global.Migrate()
	//query()
	//updateSave()

	//update()
	//updates()
	//delete()

	HighQuery()
}
