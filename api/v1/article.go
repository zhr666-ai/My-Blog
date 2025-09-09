package v1

import (
	"My-Blog/model"
	"My-Blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	id, _ := strconv.Atoi(c.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code, total := model.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个文章信息
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func GetArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	data, code, total := model.GetArt(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArt(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data) //将HTTP请求体中的JSON数据解析并绑定到指定结构体变量data上
	code = model.EditArt(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) //strconv.Atoi 函数用于将字符串转换为整数
	code := model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
