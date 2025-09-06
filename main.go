package main

import (
	"My-Blog/model"
	"My-Blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()
	routes.InitRouter()
}
