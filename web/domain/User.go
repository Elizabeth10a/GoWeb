package domain

type User struct {
	//toJson  from Form
	Name string `json:"name" form:"name"`
	Addr string `json:"addr" form:"addr"`
	Psw  string `json:"psw" form:"psw"`
	Age  int    `json:"age" form:"age"`
}
