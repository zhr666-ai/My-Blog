package model

import (
	"My-Blog/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// sql.DB 是底层基石，gorm.DB 是基于它的高层 ORM 封装。前者灵活但繁琐，后者简洁但有一定性能开销
var db *gorm.DB //创建一个 GORM 数据库会话对象
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", utils.DbUser, utils.DbPassWord, utils.DbHost, utils.DbPort, utils.DbName)
	//这里和视频中有所区分
	//&gorm.Config{} 作用：配置 GORM 的全局行为（可选参数，不填则使用默认配置）
	//可自定义日志模式、命名策略、连接池等
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败：", err)
		return
	}
	//数据迁移
	//用于根据 Go 结构体（模型）的定义，自动在数据库中创建或更新对应的表结构。
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	//用于配置数据库取得连接池
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量。
	//db.DB().SetMaxIdleConns(10)
	//
	//// SetMaxOpenConns 设置打开数据库连接的最大数量。
	//db.DB().SetMaxOpenConns(100)
	//
	//// SetConnMaxLifetime 设置了可以重新使用连接的最大时间。
	//db.DB().SetConnMaxLifetime(10 * time.Second)

	// 先获取底层sqlDB并检查错误
	//注意：db.DB() 方法的返回值是 (*sql.DB, error)，就是go标准底层数据库连接管理对象
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取连接池失败: %w", err)
		return
	}

	// 配置连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)
	//sqlDB.Close()
}
