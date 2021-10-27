package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
)

func main() {
	db := datasource.MysqlDb
	defer db.Close()

	//开启事务，后续的操作也必须使用ct，否则事务失效
	tx := db.Begin()
	//捕获panic，则rollback
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//增加
	article := Article{
		Title:   "go语言大法",
		Content: "测试内容2",
		Desc:    "测试简介2",
		Tags: []Tag{
			{
				Name: "golang",
				Desc: "go标签",
			},
			{
				Name: "c++的替代品",
				Desc: "c++标签",
			},
		},
	}
	tx.Create(&article)

	var user1 User2
	db.Preload("Articles").Find(&user1, 13)
	if err := tx.Model(&user1.Articles).Where("`desc` = ?", "演员").Update("content", "一个好演员哈哈哈哈").Error; err != nil {
		tx.Rollback()
	}
	result := tx.Commit()
	if result.Error != nil {
		tx.Rollback()
	}
}
