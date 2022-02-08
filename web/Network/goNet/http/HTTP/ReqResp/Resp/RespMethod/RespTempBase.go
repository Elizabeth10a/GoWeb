package RespMethod

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
使用Go的 Web模板引擎需要以下两个步骤:w
	(1)对文本格式的模板源进行语法分析，创建一个经过语法分析的模板结构,其中模板源既可以是一个字符串，
		也可以是模板文件中包含的内容。v=
	(2)执行经过语法分析的模板，将ResponseWriter,和模板所需的动态数据传递给模板引擎，
		被调用的模板引擎会把经过语法分析的模板和传入的数据结合起来，生成出最终的HTML，并将这些HTML传递给ResponseWriter。
*/

// Temps 模板引擎处理响应数据
type Temps struct {
	MSG, Port string
}

func (t *Temps) UseTemp1() {
	//创建处理器函数
	fmt.Println("-------UseTemp1-----")
	http.HandleFunc("/UseTemp1", func(w http.ResponseWriter, r *http.Request) {

		//解析模板文件
		filePath := "Network/Protocol/HTTP/ReqResp/templates/tempBase.html"

		//Go会创建一个新的模板,只会返回一个模板,以第一个文件的文件名作为模板的名字，
		//至于其他文件对应的模板则会被放到一个map 中。
		temp, err := template.ParseFiles(filePath)
		/*		相当与以下两句
				temp = template.New(filePath)
				temp,_ = temp.ParseFiles(filePath)*/

		if err != nil {
			panic(err)
		}

		//执行
		err1 := temp.Execute(w, t.MSG) //只能得到首个模板
		if err1 != nil {
			panic(err1)
		}

	},
	)

}
func (t *Temps) UseTemp2() {
	//创建处理器函数
	fmt.Println("-----UseTemp2----")
	http.HandleFunc("/UseTemp2", func(w http.ResponseWriter, r *http.Request) {

		//解析模板文件
		// t, _ := template.ParseFiles("index.html")
		//通过Must函数让Go帮我们自动处理异常
		filePath := "Network/Protocol/HTTP/ReqResp/templates/tempBase.html"
		filePath2 := "Network/Protocol/HTTP/ReqResp/templates/tempBase2.html"

		temp := template.Must(template.ParseFiles(filePath, filePath2))
		//执行

		//func (temp *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
		//使用名为name的 temp 关联的模板产生输出。

		//将响应数据在index2.html文件中显示 ; 文件名
		err := temp.ExecuteTemplate(w, "tempBase2.html", "我要去tempBase2.html中")
		if err != nil {
			panic(err)
		}
	},
	)

}
func (t *Temps) ListenAndServe() {
	err := http.ListenAndServe(t.Port, nil)
	if err != nil {
		panic(err)
	}
}
