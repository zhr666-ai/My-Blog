package model

import (
	"My-Blog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model //GORM 框架提供的一个内置结构体，支持软删除
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
	//db.Preload("Category")预加载模型中的 Category 关联字段
	//一次性将主表数据和关联的 Category 数据查询出来，而不是先查主表再循环查关联表。
	//Limit(pageSize)限制查询返回的最大记录数，用于分页（每页显示 pageSize 条数据）
	//Offset((pageNum-1)*pageSize)设置查询的偏移量
	//Where("cid = ?", id)添加查询条件
	//Find(&cateArtList)执行查询，并将结果存入 cateArtList 切片
	//Count(&total)统计符合查询条件的总记录数（不含 Limit 和 Offset 的限制）
	//.Error获取整个链式调用中可能出现的错误
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
	//make 是 Go 中用于初始化引用类型（map、slice、channel）的内置函数，会主动分配内存
	//map[string]interface{} 定义了 map 的类型
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	//db.Model(&art)指定操作的模型，GORM 会根据结构体类型确定对应的数据库表（默认是结构体名的复数形式）
	//作用：明确要更新的表
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
