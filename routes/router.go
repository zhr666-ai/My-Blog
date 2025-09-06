package routes

import (
	v1 "My-Blog/api/v1"
	"My-Blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//是在 Gin 框架中设置运行模式的常见写法，通常用于根据配置动态切换开发环境（debug）和生产环境（release）。
	gin.SetMode(utils.AppMode)
	r := gin.Default() //默认路由引擎，自带两个中间件
	routerV1 := r.Group("api/v1")
	{
		//用户模块的路由接口
		routerV1.POST("user/add", v1.AddUser)
		routerV1.GET("users", v1.GetUser)
		routerV1.PUT("users/:id", v1.EditUser)
		routerV1.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		routerV1.POST("category/add", v1.AddCategory)
		routerV1.GET("category", v1.GetCate)
		routerV1.PUT("category/:id", v1.EditCate)
		routerV1.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由接口
		routerV1.POST("article/add", v1.AddArticle)
		routerV1.GET("article", v1.GetArt)
		routerV1.PUT("article/:id", v1.EditArt)
		routerV1.DELETE("article/:id", v1.DeleteArt)
	}
	r.Run(utils.HttpPort)
}
