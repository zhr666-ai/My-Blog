package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	//文章分类
	Category Category `gorm:"foreignKey:Cid"` //指定外键为 Cid
	//文章的标题
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	//对应文章分类id
	Cid int `gorm:"type:int ;not null" json:"cid"`
	//文章的描述
	Desc string `gorm:"type:varchar(200)" json:"desc"`
	//文章的主题
	Content string `gorm:"type:longtext" json:"content"`
	//文章的图片
	Img string `gorm:"type:varchar(100)" json:"img"`
}
