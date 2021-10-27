package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
)

//一对一模型操作

func main() {
	db := datasource.MysqlDb
	defer db.Close()

	//增加
	userProfile := UserProfile{
		Id:    2,
		Pic:   "1.jpg",
		CPic:  "2.jpg",
		Phone: "15911118888",
		User: User{
			Name: "Jack",
			Age:  18,
			Addr: "北京 昌平 融泽嘉园",
		},
	}

	db.Create(&userProfile)

	//查询
	//第一种方式
	userProfile1 := UserProfile{}
	db.First(&userProfile1, 2)
	db.Model(&userProfile1).Association("User").Find(&userProfile1.User)
	fmt.Println(userProfile1)

	//第二种方式 Preload
	userProfile2 := UserProfile{}
	db.Preload("User").Find(&userProfile2, 2)
	fmt.Println(userProfile2)

	//第三种方式
	userProfile3 := UserProfile{}
	db.First(&userProfile3, 2)
	user := User{}
	db.Model(&userProfile3).Related(&user, "User")
	fmt.Println(userProfile3)
	fmt.Println(user)

	//更新
	userProfile4 := UserProfile{}
	db.Preload("User").First(&userProfile4, 2)
	//更新关联的表记录
	db.Model(userProfile4.User).Update("Name", "Rose")                  //单字段
	db.Model(userProfile4.User).Update(User{Age: 19, Addr: "湖北 武汉 光谷"}) //多字段

	//删除
	userProfile5 := UserProfile{}
	db.Preload("User").First(&userProfile5, 2)
	//删除本表记录
	//db.Delete(&userProfile5)
	//删除关联的表记录
	db.Debug().Delete(&userProfile5.User)

}
