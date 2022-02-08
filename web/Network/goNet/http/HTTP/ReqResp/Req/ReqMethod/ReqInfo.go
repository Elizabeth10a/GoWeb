package ReqMethod

/*
Request：用户请求的信息，用来解析用户的请求信息，包括 post、get、cookie、url 等信息
Response：服务器需要反馈给客户端的信息
Conn：用户的每次请求链接
Handler：处理请求和生成返回信息的处理逻辑
*/
import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ReqHeader struct {
	Port string
}

//获取请求头信息
func getReqHeader(w http.ResponseWriter, req *http.Request) {

	fmt.Println("请求地址：req.URL.Path		", req.URL.Path)
	fmt.Println("查询字符串：req.URL.RawQuery	", req.URL.RawQuery)
	//fmt.Println("请求头信息：req.Header		", req.Header) //map
	fmt.Println("---------------------")
	//遍历map内容
	for key, value := range req.Header {
		fmt.Println(key, value)

	}
	fmt.Println("---------------------")

	fmt.Println("请求：User-Agent	信息	", req.Header["User-Agent"])    //数组; 特定元素的值
	fmt.Println("请求：User-Agent属性值", req.Header.Get("User-Agent")) //字符串

}
func getReqBody(w http.ResponseWriter, req *http.Request) {
	fmt.Println("--- req.ContentLength--")

	//创建byte切片
	body := make([]byte, req.ContentLength) //获取请求体重内容的长度
	//将请求体中的内容读到body中
	//它返回读取的字节数 (0 <= n <= len(p)) 和遇到的任何错误,	在输入流末尾返回非零字节数的 Reader 可能返回 err == EOF 或 err == nil
	readLen, err := req.Body.Read(body)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println("readBytesLen：", readLen)
	fmt.Println("请求体中的内容有：", strings.Split((string(body)), "&"))
	//username=sdfsdfasd&password=asd&interest=football

	fmt.Println("req.FormValue:" + req.FormValue("interest"))

}

func (m *ReqHeader) GetReqHeader() {

	http.HandleFunc("/GetReqHeader", getReqHeader)

}

// GetReqBody 获取请求体信息 get没有请求体
func (m *ReqHeader) GetReqBody() {
	http.HandleFunc("/GetReqBody", getReqBody)

}
