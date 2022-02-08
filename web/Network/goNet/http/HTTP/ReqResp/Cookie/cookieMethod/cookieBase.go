package main

import (
	"fmt"
	http "net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host, r.URL.Path)
	cookie := http.Cookie{
		Name:  "ok",
		Value: "ok",
	}
	r.AddCookie(&cookie)

}
func getCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("ok")
	if err != nil {
		panic(err)
	}
	fmt.Println(cookie.Value)
	/*根据 request 选择 handler，并且进入到这个 handler 的 ServeHTTP
	  mux.handler(r).ServeHTTP(w, r)*/
}
func main() {
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/Logout", Logout)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
