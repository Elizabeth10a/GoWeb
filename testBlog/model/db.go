package model

import (
	"GoWeb/testBlog/utils"
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDb() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DB.DbUser,
		utils.DB.DbPassWord,
		utils.DB.DbHost,
		utils.DB.DbPort,
		utils.DB.DbName,
	)
	db, err := sql.Open("mysql", dsn) //判断dsn格式是否正确
	if err != nil {
		panic(err)
		return

	}
	err = db.Ping() //尝试连接数据库；验证账号密码
	if err != nil {
		panic(err)
		return
	}

	db1, err := gorm.Open(sqlite.Dialector{DriverName: "mysql", DSN: dsn})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		return
	}
	db1.AutoMigrate(&User{}, &Category{}, &Comment{}) //自动迁移 将模型与数据库同步
	sqlDB, _ := db1.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	return
}
