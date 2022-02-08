package main

import (
	"GoWeb/testBlog/model"
	"GoWeb/testBlog/routes"
)

func main() {
	routes.InitRouter()
	model.InitDb()
}
