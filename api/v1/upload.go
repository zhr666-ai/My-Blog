package v1

import (
	"My-Blog/model"
	"My-Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	//c是Gin框架的上下文对象
	//c.Request是底层HTTP请求对象
	//FormFile("file")是*http.Request的方法，用于从multipart/form-data类型的请求中获取指定键名的文件
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)

	c.JSON(200, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
