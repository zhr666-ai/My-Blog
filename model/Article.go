package model

import (
	"My-Blog/utils/errmsg"
	"gorm.io/gorm"
)

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

// 对数据库的操作

// 新增文章
func CreateArt(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64

	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtList).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_EXIT
	}
	return art, errmsg.SUCCSE
}

// 查询文章列表
func GetArt(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var total int64
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCSE, total
}

// 编辑文章
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = db.Model(&art).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
