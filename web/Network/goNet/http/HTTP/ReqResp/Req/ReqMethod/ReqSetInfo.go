package ReqMethod

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//User 结构体
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Port     string
}

// SetJsonResp  设置相应格式
func (u *User) SetJsonResp() {
	fmt.Println("------SetJsonResp---------")

	http.HandleFunc("/SetJsonResp", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-------setJsonResp：w.Header().Set(\"Content-Type\", \"application/json\")---------")
		//设置响应内容的类型
		w.Header().Set("Content-Type", "application/json")
		//创建User
		user := User{
			Port:     ":8080",
			ID:       1,
			Username: "admin",
			Password: "123456",
			Email:    "admin@atguigiu.com",
		}
		//将User转换为Json个数
		json, _ := json.Marshal(user)

		fmt.Println(string(json))

		//将json格式的数据响应给客户端
		w.Write(json)
	})
}

// ChangePage 跳转页面
func (u *User) ChangePage() {
	fmt.Println("----ChangePage----")
	http.HandleFunc("/ChangePage", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("--changePage:w.Header().Set(\"Location\", \"https://www.baidu.com\")----")
		//设置响应头中的Location属性
		w.Header().Set("Location", "https://www.baidu.com")
		//设置响应的状态码
		w.WriteHeader(302) //在w.Header().Set()下面
	})
}
