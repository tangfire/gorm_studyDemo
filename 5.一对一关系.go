package main

import (
	"gorm_studyDemo/global"
	"gorm_studyDemo/models"
)

func oneToOne() {
	// 创建用户，连带着创建用户详情
	//err := global.DB.Create(&models.UserModel{
	//	Name: "firefire",
	//	Age:  25,
	//	UserDetailModel: &models.UserDetailModel{
	//		Email: "52628216@qq.com",
	//	},
	//}).Error
	//fmt.Println(err)

	// 创建用户详情，关联用户
	//err := global.DB.Create(&models.UserDetailModel{
	//	Email:     "6666@qq.com",
	//	UserModel: &models.UserModel{ID: 15},
	//	// UserID: 15
	//}).Error
	//fmt.Println(err)

	//// 通过用户详情查用户 正查
	//var id = 15
	//var detail models.UserDetailModel
	//global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	//fmt.Println(detail.Email, detail.UserModel.Name)
	//
	//// 反查
	//var user models.UserModel
	//global.DB.Preload("UserDetailModel").First(&user, id)
	//fmt.Println(user.Name, user.UserDetailModel.Email)
	//

	// 级联删除写法一
	//var user = models.UserModel{ID: 15}
	//
	//global.DB.Model(&user).Select("UserDetailModel").Delete(&user)

	// 级联删除写法二
	//var user models.UserModel
	//global.DB.Unscoped().Take(&user,14)
	//global.DB.Unscoped().Select("UserDetailModel").Delete(&user)
	//

	// set null
	//var user models.UserModel
	//global.DB.Unscoped().Take(&user, 15)
	//global.DB.Create(&user).Association("UserDetailModel").Clear()
	//global.DB.Unscoped().Delete(&user)

	// 利用数据库外键
	var user models.UserModel
	global.DB.Unscoped().Take(&user, 16)
	global.DB.Unscoped().Delete(&user)

}

func main() {
	global.Connect()
	//global.Migrate()
	oneToOne()
}
