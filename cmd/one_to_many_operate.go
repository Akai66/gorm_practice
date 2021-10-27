package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
)

func main() {
	db := datasource.MysqlDb
	defer db.Close()

	////增加
	//user1 := User2{
	//	Name: "John",
	//	Age:  20,
	//	Addr: "广州 汕头",
	//	Articles: []Article2{
	//		{
	//			Title:   "go内存溢出分析",
	//			Content: "内容测试1",
	//			Desc:    "简介1",
	//		},
	//		{
	//			Title:   "storm分布式部署",
	//			Content: "内容测试2",
	//			Desc:    "简介2",
	//		},
	//	},
	//}
	//
	//user2 := User2{
	//	Name: "Lula",
	//	Age:  24,
	//	Addr: "湖北 武汉",
	//	Articles: []Article2{
	//		{
	//			Title:   "kafka参数调优",
	//			Content: "内容测试3",
	//			Desc:    "简介3",
	//		},
	//		{
	//			Title:   "flume服务简介",
	//			Content: "内容测试4",
	//			Desc:    "简介4",
	//		},
	//	},
	//}
	//
	//db.Create(&user1)
	//db.Create(&user2)
	//
	////查询
	////第一种方式
	//user3 := User2{}
	//db.First(&user3, 1)
	//db.Model(&user3).Association("Articles").Find(&user3.Articles)
	//fmt.Println(user3)
	//
	////第二种方式 Preload
	//user4 := User2{}
	//db.Preload("Articles").Find(&user4, 2)
	//fmt.Println(user4)
	//
	////第三种方式
	//user5 := User2{}
	//db.First(&user5, 2)
	//var articles []Article2
	//db.Model(&user5).Related(&articles, "Articles")
	//fmt.Println(articles)
	//
	////更新
	//user6 := User2{}
	//db.First(&user6, 1)
	////这里一定要指定更新条件，否则会将所有满足条件的记录全部更新，不仅仅是和该user关联的
	//db.Debug().Model(&user6.Articles).Where("content like ? and uid = ?", "%测试%", user6.Id).Update("content", "内容测试updated1")

	//删除
	//先查询
	user7 := User2{}
	db.Preload("Articles").Find(&user7, 2)
	//再删除，必须指定条件，否则会删除所有满足条件的记录
	//desc是mysql内置关键字，必须加``
	db.Debug().Where("`desc` like ? and uid = ?", "%简介%", user7.Id).Delete(&user7.Articles)
}
