package routes

import (
	"GoWeb/testBlog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.App.AppMode)
	r := gin.Default()
	r.GET("hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "msg",
		})
	})
	fmt.Println(utils.App.HttpPort)
	r.Run(utils.App.HttpPort)
}
