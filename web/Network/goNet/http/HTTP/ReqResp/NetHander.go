package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//creatServer()
	useMyHandler()
}

// MyHandler 自定义Handler
type MyHandler struct {
	path string
	port int
}

//ServeHTTP 实现ServeHTTP 方法
func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "-----自定义-MyHandler----")
}

//useMyHandler 使用自定义
func useMyHandler() {
	//------------------------
	m := MyHandler{path: "/MyHandler", port: 8000}
	//http.Handle(m.path, &m)
	//err := http.ListenAndServe(":"+strconv.Itoa(m.port), nil) //创建路由
	//if err != nil {
	//	panic(err)
	//}
	//------------------------
	mux := http.NewServeMux() //创建多路复用器
	//http.HandleFunc("/", handler) //原来用这个

	mux.HandleFunc("/", handler)
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.ListenAndServe(":"+strconv.Itoa(m.port), mux)

	//------------------------

	/*	//创建server结构体,使用其ListenAndServe() 方法
		server := http.Server{
			Addr:        ":" + strconv.Itoa(m.port),
			Handler:     &m,
			ReadTimeout: 2 * time.Second,
		}
		server.ListenAndServe()*/
	//------------------------

}

//创建简单服务器
func creatServer() {
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", handler) //回声请求调用处理程序
	//创建路由
	err := http.ListenAndServe(":8000", nil) //默认端口：http 80,https 443;handle 为nil-->默认 DefaultServeMux
	if err != nil {
		panic(err)
	}

}

//HandlerFunc type是一个适配器，通过类型转换让我们可以将普通的函数作为HTTP处理器使用。如果f是一个具有适当签名的函数，HandlerFunc(f)通过调用f实现了Hander接口。
//处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)   //当前路径;URL.Path = "/"
	fmt.Fprintf(w, "URL.Path =-= %q\n", r.URL.Path) //当前路径;URL.Path = "/"

}
