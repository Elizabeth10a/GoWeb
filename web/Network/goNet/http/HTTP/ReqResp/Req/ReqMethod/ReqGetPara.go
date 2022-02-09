package ReqMethod

import (
	"fmt"
	"net/http"
)

//获取请求参数
func getReqPare(w http.ResponseWriter, req *http.Request) {
	//解析表单，在调用r.Form之前必须执行该操作
	/*ParseForm 填充 req.Form 和 r.PostForm。
	对于所有请求，ParseForm 解析来自 URL 的原始查询并更新 r.Form。
	对于 POST、PUT 和 PATCH 请求，它还读取请求正文，将其解析为表单并将结果放入 r.PostForm 和 r.Form。请求正文参数优先于 r.Form 中的 URL 查询字符串值。
	如果请求正文的大小尚未受 MaxBytesReader 限制，则大小上限为 10MB。
	对于其他 HTTP 方法，或者当 Content-Type 不是 application/x-www-form-urlencoded 时，不读取请求正文，并且 r.PostForm 被初始化为非 nil 的空值。

	ParseMultipartForm 自动调用 ParseForm。ParseForm 是幂等的。*/

	fmt.Println("------getReqPare-------")
	/*
		r.ParseForm() //解析 url 传递的参数，对于 POST 则解析响应包的主体（request body）
		//注意:如果没有调用 ParseForm 方法，下面无法获取表单的数据*/
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	////获取请求参数

	/*如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面*/
	fmt.Println("请求参数有：", req.Form) //
	/*
		请求参数有： aaaaa; map[password:[aaaaaaaa nnnnn] username:[aaaaaa nnnn]]
		请求参数有： map[password:[aaaaaaaaaaaa] psw:[nnnnn] user:[nnnn] username:[aaaaaa]]*/

	//req.PostForm 仅仅支持enctype="application/x-www-form-urlencoded"
	fmt.Println("POST请求的form表单中的请求参数有：", req.PostForm) //不会包含url中的数据

	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值（会自动调用 req.ParseForm() 进行解析）
	fmt.Println("URL中的user请求参数的值是：", req.FormValue("user")) //和Form同名优先 Form
	fmt.Println("Form表单中的username请求参数的值是：", req.PostFormValue("username"))

}
