package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	ip := "127.0.0.1"
	port := 8080
	addr := fmt.Sprintf("%s%d", ip, port)
	listen, err := net.Listen("tcp", string(addr))
	if err != nil {
		panic("err!!")
		return

	}
	//连接
	conn, err := listen.Accept()
	defer conn.Close()
	if err != nil {
		panic("Accept err!!")
		return
	}
	fmt.Println("-- tcp 链接成功------")
	buf := make([]byte, 1024)

	write, err := conn.Write(buf)
	if err != nil {
		panic("Write err!!")

		return
	} //读取数据到buf
	fmt.Println("Client ---》Server,数据长度:", write, "数据:", string(buf))

	//数据处理
	upper := strings.ToUpper(string(buf))

	write, err = conn.Write([]byte(upper))
	if err != nil {
		return
	}
	fmt.Println("Server ---》Client ,数据长度:", write, "数据:", upper)

}
