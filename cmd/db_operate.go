package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
)

func main() {
	//常见的db操作

	db := datasource.MysqlDb
	defer db.Close()

	//启动Logger，显示详细信息，所有的查询sql日志都会打印，使用db.Debug()只显示指定操作的sql
	db.LogMode(true)

	//Pluck 查询单列，将结果扫描进切片
	var names []string
	db.Model(&User2{}).Pluck("name", &names)
	fmt.Println(names)

	//查询多列
	var users []User2
	db.Select("name,age").Find(&users)
	fmt.Println(users)

}
