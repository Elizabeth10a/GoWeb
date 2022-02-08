package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type DbConfig struct {
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
}

type AppConfig struct {
	AppMode  string
	HttpPort string
	JwtKey   string
}

/*type AppConfig struct {
	AppMode  string
	HttpPort string
	JwtKey   string
}*/

var (
	App *AppConfig

	DB *DbConfig
)

func init() {

	file, err := ini.Load("testBlog/config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	c := AppConfig{}
	App = c.AddAppConfig(file)

	d := DbConfig{}
	DB = d.AddDbConfig(file)
}

func (c *AppConfig) AddAppConfig(file *ini.File) *AppConfig {
	c.AppMode = file.Section("server").Key("AppMode").MustString("debug")
	c.HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	c.JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
	return c
}

func (c *DbConfig) AddDbConfig(file *ini.File) *DbConfig {
	c.Db = file.Section("database").Key("Db").MustString("debug")
	c.DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	c.DbPort = file.Section("database").Key("DbPort").MustString("3306")
	c.DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	c.DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	c.DbName = file.Section("database").Key("DbName").MustString("ginblog")
	return c

}
