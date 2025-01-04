package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	dsn := "root:8888.216@tcp(127.0.0.1:3306)/gorm_new_db?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("dsn格式错误 %s", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败 %s", err)
	}
	fmt.Println(db)

	//exec, err := db.Exec("insert into users(name,age,email) values('tangfire',18,'52628217@qq.com')")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(exec)

	rows, _ := db.Query("select id,name from users")
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		fmt.Println(err, id, name)
	}

	var id int
	var name string

	err = db.QueryRow("select id,name from users").Scan(&id, &name)
	fmt.Println(err, id, name)

}
