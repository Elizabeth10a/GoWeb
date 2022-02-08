package route

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"time"
)

func (rt *RouteTools) UseAllRoute(engine *gin.Engine) {

	para := getPara{}
	para.getInfo(engine)
	para.fromUrl(engine)
	para.fromForm(engine)
	para.fromUrlToObj(engine)

	sT := showTemplate{}
	sT.getHtml(engine)

	data := returnData{}
	data.getJson(engine)
	data.getXml(engine)

	//paraPost(engine)

}

// RouteTools  路由方法有 GET, POST, PUT, PATCH, DELETE 和 OPTIONS，还有Any，可匹配以上任意类型的请求。
type RouteTools struct {
	name string
}

// AllMethodMap type FuncMap map[string]interface{}
// map: string -> fun
func (rt *RouteTools) AllMethodMap() template.FuncMap {
	t := template.FuncMap{}
	t["UnixToTime"] = UnixToTime
	t["StringJoin"] = StringJoin
	return t
}

func UnixToTime(timeStamp int) string {
	unix := time.Unix(int64(timeStamp), 0)
	return unix.Format("2021-01-02 15:04:05")
}
func StringJoin(str, str2 string) string {
	return str + "------" + str2
}
