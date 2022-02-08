package main

import (
	"GoWeb/web/Network/goNet/http/HTTP/ReqResp/Req/ReqMethod"
	"net/http"
)

func main() {

	user := ReqMethod.User{
		Port:     ":8080",
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@atguigiu.com",
	}
	user.SetJsonResp()
	user.ChangePage()
	err := http.ListenAndServe(user.Port, nil)
	if err != nil {
		panic(err)
	}
}
