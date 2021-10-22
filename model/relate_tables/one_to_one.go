package relate_tables


//一对一
//外键放在任何一方都可以
//注意：gorm并不会真的在数据库中创建外键，只是模型间的关联关系


//属于关系
//外键和关系在同一方，有关系的那一方属于另外一个模型
//以下写法就是UserProfile 属于 User

//type User struct {
//	Id int
//	Name string
//	Age int
//	Addr string
//}
//
//type UserProfile struct {
//	Id int
//	Pic string
//	CPic string
//	Phone string
//	User User `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`   //指定关联关系,外键是UId字段，对应User表的Id字段
//	UId int //外键字段
//}


//包含关系
//外键和关系不在同一方，有关系的那一方包含另外一个模型
//以下写法就是UserProfile 包含 User

type User struct {
	Id int
	Name string
	Age int
	Addr string
	PId int
}

type UserProfile struct {
	Id int
	Pic string
	CPic string
	Phone string
	User User `gorm:"ForeignKey:PId;AssociationForeignKey:Id"`   //指定关联关系,外键是PId字段，对应UserProfile表的Id字段
}