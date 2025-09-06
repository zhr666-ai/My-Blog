package model

import "gorm.io/gorm"

type User struct {
	//gorm.Model 是 GORM 库中定义的一个基础模型结构体，
	//包含了一些常用的字段，用于被其他模型结构体嵌入
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username"`
	Password string `gorm:"type: varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}
