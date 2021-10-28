package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
	"time"
)

func main() {
	db := datasource.MysqlDb
	defer db.Close()

	//====DML操作====
	//create
	//userA := User{Name: "Mark",Age: 33}
	//db.Create(&userA)

	//select
	//根据主键查询
	var user User
	db.First(&user)
	var userB User
	db.First(&userB, 2)

	//随机获取一条记录
	var userT User
	db.Take(&userT)

	//根据主键查询最后一条记录
	var userL User
	db.Last(&userL)

	//查询所有的记录
	var users []User
	db.Find(&users)

	//where 条件

	db.Where("name = ?", "Mark").Find(&users)    //全等于
	db.Where("name <> ?", "Mark").Find(&users)   //不等于
	db.Where("name like ?", "%Ro%").Find(&users) //like
	lowDate := time.Now().Add(-7 * 24 * time.Hour)
	db.Where("updated_at > ?", lowDate).Find(&users)               //日期，时间
	db.Where("age > ?", 20).First(&users)                          //比较
	db.Where("name IN (?)", []string{"Jack", "Rose"}).Find(&users) //in
	db.Where("name = ? and age > ?", "Jack", 19).Find(&users)      //and or
	db.Where("age between ? and ?", 18, 25).Find(&users)           //between and
	db.Where([]int{2, 4}).Find(&users)                             //主键切片
	fmt.Println(users)

	//update
	db.Model(&user).Where("name=?", "Jack").Update("age", 20)

	//delete
	db.Delete(&user, 1) //如果组合了gorm.Model，不会物理删除，只修改deleted_at字段
	db.Where("name = ?", "Rose").Delete(&user)
}
