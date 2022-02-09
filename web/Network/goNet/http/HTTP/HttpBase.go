package main

//HTTP超文本传输协议(HTTP-Hypertext transfer protocol)，是一个属于应用层的面向对象的协议，由于其简捷、快速的方式，适用于分布式超媒体信息系统。
//它于1990年提出，经过几年的使用与发展，得到不断地完善和扩展。
//它是一种详细规定了浏览器和万维网服务器之间互相通信的规则，通过因特网传送万维网文档的数据传送协议。
/*
	传输的内容 ：报文
	客户端-->服务端：请求
	服务端-->客户端：响应
	HTTP1.0：每次请求都需要新建一连接
	HTTP1.0：http 请求链接复用
	HTTP2.0：https
*/
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpPostGet(method, url, contentType, body string) {
	resp, err := http.Get(url)

	switch method {
	case http.MethodGet:
		fmt.Println(http.MethodGet)

		resp, err = http.Get(url)
	case http.MethodPost:
		//body io.Reader
		fmt.Println(http.MethodPost)
		resp, err = http.Post(url, contentType, strings.NewReader(body))
	default:
		resp, err = http.Get(url)

	}

	if err != nil {
		panic(err)
	}

	defer func() { resp.Body.Close() }()
	showInfo(resp)
}
func httpMethod(method, url string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer func() { resp.Body.Close() }()
	showInfo(resp)
}

func showInfo(resp *http.Response) {
	ra, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(ra))

}
func main() {
	url := "http://httpbin.org/get"
	//url1 := "http://httpbin.org/put"
	//url2 := "http://httpbin.org/delete"
	//url3 := "http://httpbin.org/post"
	httpPostGet(http.MethodGet, url, "", "")

}
