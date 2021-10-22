package relate_tables

//一对多
//外键在多的一方

//下面为包含关系，User2 包含 Article

type User2 struct {
	Id int
	Name string
	Age int
	Addr string
	Articles []Article2 `gorm:"ForeignKey:UId;AssociationForeignKey:Id"`
}

type Article2 struct {
	Id int
	Title string
	Content string
	Desc string
	UId int
}
