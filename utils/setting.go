package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	//该方法会读取指定路径（config/config.ini）的 INI 格式文件
	//并返回一个 *ini.File 类型的对象，后续可通过该对象获取具体配置值。
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误：", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadQiniu(file)
}
func LoadServer(file *ini.File) {
	//从配置文件读取参数，Section是段落的意思，用于定位配置文件的某一块内容（比如server），
	//MustString 表示 “必须返回一个字符串类型的值”,如果配置文件的 [server] 段落中存在 AppMode 键，则返回它的值（比如例子中的 production）；
	//如果不存在这个键（或配置文件缺失、段落错误），则返回默认值 debug。
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("My-blog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("My-blog")

}

func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}
