package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/model"
	"log"
)

const (
	_mysqlConfig = "root:123456@(localhost:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local"
)


func main()  {
	db,err := gorm.Open("mysql",_mysqlConfig)
	if err != nil {
		log.Fatal("connect mysql: %v",err)
	}
	defer db.Close()

	//====DDL操作====

	//创建表
	//db.Table("user").CreateTable(&model.User{})  //指定表名建表
	//db.CreateTable(&model.User{})  //不指定表名，默认是复数形式，即users
	//db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").CreateTable(&model.User{})

	//删除表
	//db.DropTableIfExists(&model.User{})

	//检查表是否存在
	hasUserTable := db.HasTable(&model.User{})
	log.Printf("hasUserTable: %t",hasUserTable)

	////添加唯一索引
	//db.Model(&model.User{}).AddUniqueIndex("uniq_idx_name_age","name","age")
	//
	////自动迁移，根据模型同步表的最新结构，自动迁移仅仅会创建表，添加缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据
	//db.AutoMigrate(&model.User{})


	//====DML操作====
	//增
	//db.Create(&model.User{Id: 1,Name: "Jack",Age: 18})

	//查
	var user model.User
	//db.First(&user,1) //默认使用id字段
	//log.Printf("user: %v",user)

	db.First(&user,"id=?",1) //查询指定字段
	log.Printf("user: %v",user)

	//改
	db.Model(&user).Update("age",22).Update("name","Jack")
	log.Printf("user: %v",user)  //user对象和数据库记录均会被修改

	//删
	db.Delete(&user)

}
