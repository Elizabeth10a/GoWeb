package main

import (
	"GoWeb/web/gin/route"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default() //返回默认路由引擎
	tools := route.RouteTools{}
	//加载模板函数
	engine.SetFuncMap(tools.AllMethodMap())

	tools.UseAllRoute(engine)

	//加载模板 多目录操作-
	tempPath := "web/gin/templates/**/*"
	engine.LoadHTMLGlob(tempPath)

	//配置静态web目录（配置文件）
	staticPath := "web/gin/static"

	engine.Static("/static", staticPath)
	engine.Run(":8080")
}
