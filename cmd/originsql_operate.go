package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm_practice/datasource"
	. "gorm_practice/model/relate_tables"
)

func main() {
	//原生sql操作

	db := datasource.MysqlDb
	defer db.Close()

	//增加
	db.Exec("insert into article2 (title,content,`desc`,uid) values (?,?,?,?)", "kafka深入浅出", "内容", "kafka消息队列", 1)

	//查询
	var articles []Article2
	db.Raw("select * from article2 where uid = ?", 1).Find(&articles)
	fmt.Println(articles)

	//更新
	db.Exec("update article2 set content = ? where title like ?", "kafka详解....", "%kafka%")

	//删除
	db.Exec("delete from article2 where id = ?", 27)

}
