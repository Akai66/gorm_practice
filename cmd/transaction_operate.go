package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "gorm_practice/model/relate_tables"
	"gorm_practice/myconst"
	"log"
)

func main() {
	db, err := gorm.Open("mysql", myconst.MysqlConfig)
	if err != nil {
		log.Fatalf("connect mysql: %v", err)
	}
	defer db.Close()

	db.LogMode(true)

	//开启事务，后续的操作也必须使用ct，否则事务失效
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//user := User2{
	//	Name: "ChengLong",
	//	Age:  53,
	//	Addr: "香港 九龙",
	//	Articles: []Article2{
	//		{
	//			Title:   "演员的职业素养",
	//			Content: "内容测试1",
	//			Desc:    "演员",
	//		},
	//		{
	//			Title:   "十二生肖",
	//			Content: "内容测试2",
	//			Desc:    "生肖",
	//		},
	//	},
	//}
	//tx.Create(&user)

	var user1 User2
	db.Preload("Articles").Find(&user1, 14)
	if err := tx.Model(&user1.Articles).Where("`desc` = ?", "演员").Update("content", "一个好演员啊").Error; err != nil {
		tx.Rollback()
	}
	result := tx.Commit()
	if result.Error != nil {
		tx.Rollback()
	}
}
