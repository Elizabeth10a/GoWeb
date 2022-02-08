package route

import (
	"GoWeb/web/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

//gin get 方法

type returnData struct {
}

//获取json数据
func (rd *returnData) getJson(engine *gin.Engine) {
	engine.GET("getJson/json", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"name": "as",
			"age":  12,
		})
	})
	engine.GET("getJson/json1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": "a2s",
			"age":  1234,
		})
	})

	engine.GET("getJson/json2", func(context *gin.Context) {
		user := &domain.User{Name: "As", Addr: "鞍山市"}
		context.JSON(200, &user)
	})

	engine.GET("getJson/jsonp", func(context *gin.Context) {
		user := &domain.User{Name: "As", Addr: "jsonp"}

		//可以读取url中的参数
		context.JSONP(200, &user)
	})

}
func (rd *returnData) getXml(engine *gin.Engine) {
	engine.GET("/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{
			"user": &domain.User{Name: "Ad", Addr: "DDs"},
		})
	})
}
