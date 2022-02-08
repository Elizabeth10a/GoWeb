package main

import (
	"GoWeb/web/Network/goNet/http/HTTP/ReqResp/Resp/RespMethod"
)

func main() {
	temp := RespMethod.Temps{
		MSG:  "发外狂1ddsd徒",
		Port: ":8080",
	}
	//temp.UseTemp1()
	temp.UseTemp2()
	temp.ListenAndServe()
}
