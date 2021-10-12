package model

type User struct {
	Id int
	Name string
	Age uint8
}

func NewUser(Id int,Name string,Age uint8) *User {
	return &User{Id,Name,Age}
}