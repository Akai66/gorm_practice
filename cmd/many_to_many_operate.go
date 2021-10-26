package main

import (
	"fmt"
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

	//增加
	article := Article{
		Title:   "测试标题1",
		Content: "测试内容1",
		Desc:    "测试简介1",
		Tags: []Tag{
			{
				Name: "python",
				Desc: "python标签",
			},
			{
				Name: "django",
				Desc: "django标签",
			},
		},
	}
	db.Create(&article)

	//查询
	article2 := Article{}
	db.Preload("Tags").Find(&article2, 1)
	fmt.Println(article2)

	//更新
	article3 := Article{}
	db.Preload("Tags").Find(&article3, 2)
	db.Model(&article3.Tags).Where("name = ?", "python").Update("name", "人工智能")

	//删除
	article4 := Article{}
	db.Preload("Tags").Find(&article4, 1)
	db.Debug().Where("name = ?", "人工智能").Delete(&article4.Tags)
}
