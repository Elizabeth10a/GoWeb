package tools

import (
	"GoWeb/web/domain"
	"container/list"
	"database/sql"
	"fmt"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

type MysqlTool struct {
	DriverName string
	ConnUrl    string
	db         *sql.DB
	//"eliza:eliza@tcp(192.168.20.130:3306)/Test"
}

/*调用 db.Query 执行 SQL 语句, 此方法会返回一个 Rows 作为查询的结果
通过 rows.Next() 迭代查询数据.
通过 rows.Scan() 读取每一行的值
调用 db.Close() 关闭查询*/
func init() {
	fmt.Println("---MysqlTool---")
}

func (t MysqlTool) ChangeInfo(mt MysqlTool) {
	t = mt
}
func (t *MysqlTool) GetConn() *sql.DB {
	t.db, _ = sql.Open(t.DriverName, t.ConnUrl)

	fmt.Println("---数据库链接成功---")
	return t.db
}

/*调用 db.Query 执行 SQL 语句, 此方法会返回一个 Rows 作为查询的结果
通过 rows.Next() 迭代查询数据.
通过 rows.Scan() 读取每一行的值
调用 db.Close() 关闭查询*/

func (t *MysqlTool) Add(name string, age int) {
	ret, err := t.db.Exec(`INSERT INTO User (name, age) VALUES ("xys", 23)`)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}

	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}
func (t *MysqlTool) FindAll() *list.List {
	rows, err := t.db.Query(`select  id,name, age from User`)
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return nil
	}
	l := list.New()
	user := domain.User{}
	var id int
	for rows.Next() {
		rows.Scan(&id, &user.Name, &user.Age)
		fmt.Println(id, user)
		l.PushBack(&user)
	}
	return l
}

func (t *MysqlTool) FindOneByID(id int) {

	var name string
	var age int
	err := t.db.QueryRow("select name,age from User where id = ?", id).Scan(&name, &age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(domain.User{
		Name: name,
		Age:  age,
	})
}

func (t *MysqlTool) InsertData(name string, age int) {
	stmt, _ := t.db.Prepare(`INSERT INTO User (name, age) VALUES (?, ?)`)
	defer stmt.Close()

	ret, err := stmt.Exec(name, age)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (t *MysqlTool) GetDataById(id int) string {
	stmt, err := t.db.Prepare("SELECT * FROM User WHERE id = ?")
	if err != nil {
		log.Fatalln(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	defer rows.Close()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return "err"
	}
	for rows.Next() {
		user := domain.User{}
		_ = rows.Scan(&user.Name, &user.Age)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		fmt.Println("get data:", &user)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
	return ""
}
