package main

import "gorm_studyDemo/global"

func oneToOne() {
	// 创建用户，连带着创建用户详情

}

func main() {
	global.Connect()
	global.Migrate()
}
