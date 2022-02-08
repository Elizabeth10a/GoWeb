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
)

func get(url string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)

	}
	defer func() { _ = resp.Body.Close() }()
	ra, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(ra))

}
func post(url string) {
	resp, err := http.Post(url, "", nil)

	if err != nil {
		panic(err)

	}
	defer func() { _ = resp.Body.Close() }()
	ra, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(ra))

}

func put(url string) {
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
		panic(err)
	}
	do, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer func() { _ = do.Body.Close() }()
	ra, err := ioutil.ReadAll(do.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(ra))

}

func del(url string) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	do, err := http.DefaultClient.Do(req)

	if err != nil {
		return
	}

	defer func() { _ = do.Body.Close() }()
	ra, err := ioutil.ReadAll(do.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(ra))

}
func main() {
	url := "http://httpbin.org/get"
	url1 := "http://httpbin.org/put"
	url2 := "http://httpbin.org/delete"
	url3 := "http://httpbin.org/post"
	//get(url)
	//post(url)
	get(url)
	/*  "args": {},
	  "headers": {
	    "Accept-Encoding": "gzip",
	    "Host": "httpbin.org",
	    "User-Agent": "Go-http-client/1.1",
	    "X-Amzn-Trace-Id": "Root=1-617aaa50-29227db93a3b0ea34f25d160"
	  },
	  "origin": "222.88.64.189",
	  "url": "http://httpbin.org/get"
	}*/
	put(url1)
	/*{
	  "args": {},
	  "data": "",
	  "files": {},
	  "form": {},
	  "headers": {
	    "Accept-Encoding": "gzip",
	    "Content-Length": "0",
	    "Host": "httpbin.org",
	    "User-Agent": "Go-http-client/1.1",
	    "X-Amzn-Trace-Id": "Root=1-617aaa50-0ce5239e5ed44a4953e26051"
	  },
	  "json": null,
	  "origin": "222.88.64.189",
	  "url": "http://httpbin.org/put"
	}*/
	del(url2)
	/*{
	  "args": {},
	  "data": "",
	  "files": {},
	  "form": {},
	  "headers": {
	    "Accept-Encoding": "gzip",
	    "Host": "httpbin.org",
	    "User-Agent": "Go-http-client/1.1",
	    "X-Amzn-Trace-Id": "Root=1-617aaa50-113214a62ca4fb7d7e15d0f9"
	  },
	  "json": null,
	  "origin": "222.88.64.189",
	  "url": "http://httpbin.org/delete"
	}*/
	post(url3)
	/*{
	    "args": {},
	    "data": "",
	    "files": {},
	    "form": {},
	    "headers": {
	      "Accept-Encoding": "gzip",
	      "Content-Length": "0",
	      "Host": "httpbin.org",
	      "User-Agent": "Go-http-client/1.1",
	      "X-Amzn-Trace-Id": "Root=1-617aaa52-66e6d76454d3978542283dca"
	    },
	    "json": null,
	    "origin": "222.88.64.189",
	    "url": "http://httpbin.org/post"
	  }
	*/
}
