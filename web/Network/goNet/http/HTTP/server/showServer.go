package server

import (
	"fmt"
	"net/http"
)

type ReqHeader struct {
	urlList    []string
	handleList []func(w http.ResponseWriter, req *http.Request)
	Port       string
}

func (m *ReqHeader) runHandle() {
	for i := 0; i < len(m.urlList); i++ {
		fmt.Println(i, ":", m.urlList[i])
		http.HandleFunc(m.urlList[i], m.handleList[i])

	}

}
