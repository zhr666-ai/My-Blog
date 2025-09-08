package v1

import (
	"My-Blog/model"
	"My-Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, code := model.UploadFile(file, fileSize)

	c.JSON(200, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
