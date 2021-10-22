package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "gorm_practice/model/relate_tables"
	"gorm_practice/myconst"
	"log"
)

func main()  {
	db,err := gorm.Open("mysql",myconst.MysqlConfig)
	if err != nil {
		log.Fatalf("connect mysql: %v", err)
	}
	defer db.Close()

	//====DDL操作====
	//检查表是否存在
	hasUserTable := db.HasTable(&User{})
	if !hasUserTable {
		//创建表
		db.Table("user").CreateTable(&User{})  //指定表名建表
		db.CreateTable(&User{})  //不指定表名，默认是复数形式，即users
		db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").CreateTable(&User{})
	}

	//删除表
	//db.DropTableIfExists(&User{})

	//添加唯一索引
	//db.Model(&User{}).AddUniqueIndex("uniq_idx_name_age","name","age")

	//自动迁移，将模型中定义的结构自动同步到数据库，自动迁移仅仅会创建表，添加缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").AutoMigrate(
		&User{},
		&UserProfile{},
		&User2{},
		&Article2{},
		&Article{},
		&Tag{})
}
