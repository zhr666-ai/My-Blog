package model

import (
	"My-Blog/utils/errmsg"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	//gorm.Model 是 GORM 库中定义的一个基础模型结构体，
	//包含了一些常用的字段，用于被其他模型结构体嵌入
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type: varchar(20);not null " json:"password" validate:"required,min=6,max=20"label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role"validate:"required,gte=2"label:"角色码"`
}

// 对数据库的操作
// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPwd(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 钩子函数
// 注意方法和钩子函数的区别，方法得需要手动调用，跟钩子函数不需要手动调用
func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Password = ScryptPwd(u.Password)
	return nil
}

// 密码加密
func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	Fpwd := base64.StdEncoding.EncodeToString(HashPw)
	return Fpwd
}

// 登录验证
func CheckLogin(username, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPwd(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
