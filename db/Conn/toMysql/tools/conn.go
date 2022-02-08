package tools

import (
	"GoWeb/web/domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "eliza:eliza@tcp(192.168.20.129:3306)/GO?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.,&mysql.Config{})
	/*
		[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		username:password@protocol(address)/dbname?param=value*/

	db, err := sql.Open("mysql", "eliza:eliza@tcp(192.168.20.130:3306)/Test")
	if err != nil {
		panic(err)
	}
	results, err := db.Query("SELECT * FROM Test.User")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	user := domain.User{Name: "sd", Age: 0}
	for results.Next() {
		//fmt.Println(results.Columns())

		err = results.Scan(&user)
		fmt.Println(user.Name, user.Age)
	}
}
