package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"aaaaa"`
	Age  int    `json:"bbbb"`
}

func main() {

	user := User{
		"Reds",
		12,
	}
	b, _ := json.Marshal(user)
	fmt.Println(string(b))
	fmt.Println(json.Unmarshal(b, &user))

}
