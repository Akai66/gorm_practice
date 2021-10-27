package datasource

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
)

var MysqlDb *gorm.DB

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Logomode bool   `json:"logo_mode"`
}

func LoadMysqlConf() *MysqlConf {
	conf := MysqlConf{}
	file, err := os.Open("conf/mysql.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return &conf
}

func init() {
	conf := LoadMysqlConf()
	confStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	var err error
	MysqlDb, err = gorm.Open("mysql", confStr)
	if err != nil {
		log.Fatalln(err)
	}
	if conf.Logomode {
		MysqlDb.LogMode(true) //打开debug
	}
	MysqlDb.DB().SetMaxOpenConns(100) //最大连接数
	MysqlDb.DB().SetMaxIdleConns(50)  //最大空闲数
}
