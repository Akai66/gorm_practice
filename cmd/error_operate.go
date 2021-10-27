package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
)

func main() {
	//错误处理
	db := datasource.MysqlDb
	defer db.Close()

	//处理单个错误
	var articles []Article
	result := db.Where("desc like ?", "%测试%").Find(&articles)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	//处理多个错误
	//var user User
	var article Article
	errors := db.Where("desc = ?", "测试").Find(&article).GetErrors()
	for _, err := range errors {
		fmt.Println(gorm.IsRecordNotFoundError(err))
		fmt.Println(err)
	}

}
