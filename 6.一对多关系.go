package main

import (
	"gorm_studyDemo/global"
	"gorm_studyDemo/models"
)

func oneToMany() {
	// 来了一个女神，自带俩舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "xiaomei",
	//	BoyList: []models.BoyModel{
	//		{Name: "one"},
	//		{Name: "two"},
	//	},
	//}).Error
	//fmt.Println(err)

	// 来了一个女神，选了一个舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "xiaomeili",
	//	BoyList: []models.BoyModel{
	//		{ID: 2},
	//	},
	//}).Error
	//fmt.Println(err)

	// 来了一个舔狗，选了一个女神
	//err := global.DB.Create(&models.BoyModel{
	//	Name:   "three",
	//	GirlID: 2,
	//}).Error
	//fmt.Println(err)

	// 查询xiaomei的舔狗
	//var girl models.GirlModel
	//global.DB.Preload("BoyList").Take(&girl, "name = ?", "xiaomei")
	//fmt.Println(girl.Name, girl.BoyList, len(girl.BoyList))

	//var girl models.GirlModel
	//global.DB.Preload("BoyList", "name = ?", "one").Take(&girl, "name = ?", "xiaomei")
	//fmt.Println(girl.Name, girl.BoyList, len(girl.BoyList))

	// 专门查关联
	//var girl models.GirlModel
	//global.DB.Take(&girl, "name = ?", "xiaomei")
	//var boyList []models.BoyModel
	//
	//global.DB.Model(&girl).Association("BoyList").Find(&boyList)
	//fmt.Println(boyList)

	//var girl models.GirlModel
	//global.DB.Take(&girl, "name = ?", "xiaomei")
	//global.DB.Model(&girl).Association("BoyList").Replace([]models.BoyModel{
	//	{ID: 1},
	//})

	//var girl models.GirlModel
	//global.DB.Take(&girl, "name = ?", "xiaomei")
	//global.DB.Model(&girl).Association("BoyList").Append([]models.BoyModel{
	//	{ID: 2},
	//	{ID: 3},
	//})

	var girl models.GirlModel
	global.DB.Take(&girl, "name = ?", "xiaomei")
	global.DB.Model(&girl).Association("BoyList").Delete([]models.BoyModel{
		{ID: 3},
	})

}

func main() {
	global.Connect()
	//global.Migrate()
	oneToMany()

}
