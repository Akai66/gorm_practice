package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model     //gorm.Model，包括字段ID，CreatedAt，UpdatedAt，DeletedAt
	ID         int `gorm:"primary_key;AUTO_INCREMENT;not null"`
	Name       string
	Age        uint8
}

// TableName 自定义模型表名
func (user *User) TableName() string {
	return "user"
}
