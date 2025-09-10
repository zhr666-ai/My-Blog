package validator

import (
	"My-Blog/utils/errmsg"
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"                         //Go 语言常用的高性能数据验证库
	zhTrans "github.com/go-playground/validator/v10/translations/zh" //提供中文本地化支持，用于将验证错误信息翻译成中文
	"reflect"                                                        //用于反射，获取结构体标签，如label标签
)

func Validate(data interface{}) (string, int) {
	validate := validator.New() //创建一个新的validator.Validate 实例，这是验证操作的核心对象，用于定义验证规则和执行验证
	//初始化中文本地化器
	uni := unTrans.New(zh_Hans_CN.New())
	//获取中文翻译器实例
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	//将验证器与中文翻译器绑定（是错误信息自动转换成中文）
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	//自定义字段名映射（通过label标签）
	//这里的逻辑是如果结构体字段定义了label标签，则错误信息使用label的值。否则使用默认值
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	//执行验证并处理错误
	err = validate.Struct(data) //对data（必须是结构体或结构体指针）执行验证，验证规则通过结构体字段的validate标签定义
	//若验证失败，err 是 validator.ValidationErrors 类型（包含所有验证失败的字段信息）
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			//将错误信息翻译成中文并返回（只返回第一个错误）
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
