package main

import (
	"github.com/jinzhu/gorm"
	"gorm_practice/model"
	"log"
)

const (
	_mysqlConfig = "root:123456@[tcp](localhost:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local"
)


func main()  {
	db,err := gorm.Open("mysql",_mysqlConfig)
	if err != nil {
		log.Fatal("connect mysql: %v",err)
	}
	defer db.Close()

	//创建表
	db.Table("user").CreateTable(&model.User{})  //指定表名建表
	//db.CreateTable(&model.User{})  //不指定表名，默认是复数形式，即users

	//删除表
	db.DropTableIfExists(&model.User{})

	//检查表是否存在
	db.HasTable(&model.User{})

	//添加唯一索引
	db.Model(&model.User{}).AddUniqueIndex("uniq_idx_name_age","name","age")

	//自动迁移，根据模型同步表的最新结构，自动迁移仅仅会创建表，添加缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据
	db.AutoMigrate(&model.User{})

}
