package relate_tables

//多对多

//会自动创建名为Articles2Tags的中间表
/*
CREATE TABLE `Articles2Tags` (
  `article_a_id` int(11) NOT NULL,
  `tag_t_id` int(11) NOT NULL,
  PRIMARY KEY (`article_a_id`,`tag_t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8
 */


type Article struct {
	AId int `gorm:"primary_key:true"`
	Title string
	Content string
	Desc string
	Tags []Tag `gorm:"many2many:Articles2Tags;ForeignKey:AId;AssociationForeignKey:TId"`
}

type Tag struct {
	TId int `gorm:"primary_key:true"`
	Name string
	Desc string
}
