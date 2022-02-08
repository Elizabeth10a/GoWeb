package main

import (
	"GoWeb/web/Network/goNet/http/HTTP/ReqResp/Req/ReqMethod"
	"net/http"
)

func main() {
	header := ReqMethod.ReqHeader{
		Port: ":8080",
	}

	header.GetReqHeader()
	header.GetReqBody()
	header.GetReqPare()
	err := http.ListenAndServe(header.Port, nil)
	if err != nil {
		panic(err)
	}
}
