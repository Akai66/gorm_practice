package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/model/relate_tables"
	"log"
)

const (
	_mysqlConfig = "root:123456@(localhost:3306)/test_gorm?charset=utf8&parseTime=True&loc=Local"
)


func main()  {
	db,err := gorm.Open("mysql",_mysqlConfig)
	if err != nil {
		log.Fatalf("connect mysql: %v", err)
	}
	defer db.Close()
	//
	////====DDL操作====
	////检查表是否存在
	//hasUserTable := db.HasTable(&User{})
	//if !hasUserTable {
	//	//创建表
	//	db.Table("user").CreateTable(&User{})  //指定表名建表
	//	db.CreateTable(&User{})  //不指定表名，默认是复数形式，即users
	//	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").CreateTable(&User{})
	//}
	//
	////删除表
	////db.DropTableIfExists(&User{})
	//
	////添加唯一索引
	////db.Model(&User{}).AddUniqueIndex("uniq_idx_name_age","name","age")
	//
	////自动迁移，将模型中定义的结构自动同步到数据库，自动迁移仅仅会创建表，添加缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据
	//db.AutoMigrate(&User{})
	//
	//
	////====DML操作====
	////create
	////userA := User{Name: "Mark",Age: 33}
	////db.Create(&userA)
	//
	////select
	////根据主键查询
	//var user User
	//db.First(&user)
	//var userB User
	//db.First(&userB,2)
	//
	////随机获取一条记录
	//var userT User
	//db.Take(&userT)
	//
	////根据主键查询最后一条记录
	//var userL User
	//db.Last(&userL)
	//
	////查询所有的记录
	//var users []User
	//db.Find(&users)
	//
	////where 条件
	//
	//db.Where("name = ?","Mark").Find(&users)  //全等于
	//db.Where("name <> ?","Mark").Find(&users) //不等于
	//db.Where("name like ?","%Ro%").Find(&users) //like
	//lowDate := time.Now().Add(-7 * 24 * time.Hour)
	//db.Where("updated_at > ?",lowDate).Find(&users)  //日期，时间
	//db.Where("age > ?",20).First(&users) //比较
	//db.Where("name IN (?)",[]string{"Jack","Rose"}).Find(&users) //in
	//db.Where("name = ? and age > ?","Jack",19).Find(&users) //and or
	//db.Where("age between ? and ?",18,25).Find(&users) //between and
	//db.Where([]int{2,4}).Find(&users) //主键切片
	//fmt.Println(users)
	//
	//
	//
	//
	//
	//
	////update
	////db.Model(&user).Update("age",20).Where("name=?","Jack")
	//
	////delete
	////db.Delete(&user)   //如果组合了gorm.Model，不会物理删除，只修改deleted_at字段
	db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8;").AutoMigrate(
		&relate_tables.User{},
		&relate_tables.UserProfile{},
		&relate_tables.User2{},
		&relate_tables.Article2{},
		&relate_tables.Article{},
		&relate_tables.Tag{})
}
