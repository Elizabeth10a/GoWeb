package route

import (
	"GoWeb/web/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//engine.LoadHTMLGlob(tempPath)
//tempPath := "web/gin/templates"

type showTemplate struct {
}

func (st *showTemplate) getHtml(engine *gin.Engine) {

	//加载模板目录中的文件 ，文件全称
	engine.GET("/", func(context *gin.Context) {
		//"web/gin/templates"
		context.HTML(http.StatusOK, "index.html", gin.H{
			"user": &domain.User{Name: "font news数据", Addr: "DDs"},
		})
	})

	engine.GET("/back/goods", func(context *gin.Context) {
		//db/Gin/templates/goods.html
		context.HTML(http.StatusOK, "back/goods.html", gin.H{
			"user":    &domain.User{Name: "back goods数据", Addr: "DDs"},
			"names":   []string{"a", "b", "c"},
			"nilList": []string{"a", "b", "c", "d"},
			"age":     []string{"a", "b", "c", "d"},
			"date":    int(time.Now().Unix()), //当前时间的时间戳
		})
	})
	engine.GET("/font/news", func(context *gin.Context) {
		//db/Gin/templates/goods.html
		context.HTML(http.StatusOK, "font/news.html", gin.H{
			"user": &domain.User{Name: "font news数据", Addr: "/font/news"},
		})
	})
	engine.GET("/font/head", func(context *gin.Context) {
		//db/Gin/templates/goods.html
		context.HTML(http.StatusOK, "font/head.html", gin.H{
			"us":   &domain.User{Name: "font news数据", Addr: "阿斯爱上的顿"},
			"mode": "爱上是否",
			"date": 1334086974,
		})
	})

}
