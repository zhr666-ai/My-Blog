package routes

import (
	v1 "My-Blog/api/v1"
	"My-Blog/middleware"
	"My-Blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//是在 Gin 框架中设置运行模式的常见写法，通常用于根据配置动态切换开发环境（debug）和生产环境（release）。
	gin.SetMode(utils.AppMode)
	//r := gin.Default() //默认路由引擎，自带两个中间件
	r := gin.New() //初始化一个不包含任何中间件的纯净路由引擎，返回一个路由引擎实例
	r.Use(middleware.Logger())
	//gin.Recovery()中间件的作用是捕获请求处理链中的panic，恢复程序执行，返回统一的错误响应
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("users/:id", v1.EditUser) // :id 是路由参数占位符，: 的作用是标记该位置为一个动态参数，用于匹配 URL 中对应位置的具体值。
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)

		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)

		auth.PUT("article/:id", v1.EditArt)
		auth.DELETE("article/:id", v1.DeleteArt)
		//上传文件
		auth.POST("upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUser)
		router.GET("category/list", v1.GetCate)
		router.GET("article", v1.GetArt)
		router.GET("article/list/:id", v1.GetCateArt)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.POST("login", v1.Login)
	}
	_ = r.Run(utils.HttpPort)
}
