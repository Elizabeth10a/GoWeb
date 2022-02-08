package main

import (
	"GoWeb/src/Conn/toMysql/tools"
	"fmt"
)

func main() {
	var tool = tools.MysqlTool{
		DriverName: "mysql",
		ConnUrl:    "eliza:eliza@tcp(192.168.20.130:3306)/Test",
	}
	fmt.Println(tool)

	tool.GetConn()

	//tool.InsertData("as", 12)
	//fmt.Println(tool.GetDataById(2))

	tool.FindOneByID(12)
	//tool.FindAll()

}
