package model

import (
	"My-Blog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primary_key;AUTO_INCREMENT"json:"id"`
	Name string `gorm:"type:varchar(20);not null " json:"name"`
}

// 对数据库的操作
// 查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// 新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// 查询分类下的所有文章

// 查询分类列表
func GetCate(pageSize int, pageNum int) []Category {
	var cate []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

// 编辑分类
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCate(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
