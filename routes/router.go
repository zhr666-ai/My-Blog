package routes

import (
	"My-Blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//是在 Gin 框架中设置运行模式的常见写法，通常用于根据配置动态切换开发环境（debug）和生产环境（release）。
	gin.SetMode(utils.AppMode)
	r := gin.Default() //默认路由引擎，自带两个中间件
	router := r.Group("api.v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "ok",
			})
		})
	}
	r.Run(utils.HttpPort)
}
