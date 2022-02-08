package route

import (
	"GoWeb/web/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
从网页获取 数据
*/
//getPara

type getPara struct {
}

func (gp *getPara) getInfo(engine *gin.Engine) {
	engine.GET("/para/para", func(context *gin.Context) {
		context.HTML(http.StatusOK, "para/para.html", gin.H{})
	})

}

// fromUrl
//从URL 中获取值 http://localhost:8080/para?name=as
func (gp *getPara) fromUrl(engine *gin.Engine) {
	engine.GET("/getPara", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"name": context.Query("name"),
			"psw":  context.DefaultQuery("psw", "mima"),
		})
	})
}

//从 url 中获取数据 并直接转成 对象
func (gp *getPara) fromUrlToObj(engine *gin.Engine) {
	//将url 数据绑定到结构体上面 ，属性后需要加 form
	engine.GET("/getPara/getUser", func(context *gin.Context) {
		u := &domain.User{}
		if err := context.ShouldBind(u); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, u)
			fmt.Printf("%#v", u)

		}
	})
}

// 从表单中获取数据
func (gp *getPara) fromForm(engine *gin.Engine) {
	engine.POST("/getPara/fromForm1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": context.PostForm("name"),
			"pwd":  context.PostForm("pwd"),
			"age":  context.DefaultPostForm("age", "0"),
		})
	})

	engine.POST("/getPara/fromForm2", func(context *gin.Context) {
		u := &domain.User{}
		if err := context.ShouldBind(&u); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		} else {
			context.JSON(http.StatusOK, u)
			fmt.Println("转换成功")
			fmt.Printf("%#v", u)

		}
	})

}
