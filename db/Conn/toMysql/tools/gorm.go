package tools

//
//import (
//	"crypto/rsa"
//	"crypto/tls"
//	"fmt"
//	_ "github.com/go-sql-driver/mysql"
//	"gorm.io/gorm"
//	"time"
//)
//
//type User struct {
//	Name string `gorm:"not null"`
//	Age  int
//}
//
////https://gorm.io/docs/
//func main() {
//
//	type Config struct {
//		User             string            // Username
//		Passwd           string            // Password (requires User)
//		Net              string            // Network type
//		Addr             string            // Network address (requires Net)
//		DBName           string            // Database name
//		Params           map[string]string // Connection parameters
//		Collation        string            // Connection collation
//		Loc              *time.Location    // Location for time.Time values
//		MaxAllowedPacket int               // Max packet size allowed
//		ServerPubKey     string            // Server public key name
//		pubKey           *rsa.PublicKey    // Server public key
//		TLSConfig        string            // TLS configuration name
//		tls              *tls.Config       // TLS configuration
//		Timeout          time.Duration     // Dial timeout
//		ReadTimeout      time.Duration     // I/O read timeout
//		WriteTimeout     time.Duration     // I/O write timeout
//
//	}
//	db, err := gorm.Open("mysql", "eliza:eliza@tcp(192.168.20.128:3306)/Test")
//
//	if err != nil {
//		fmt.Println("连接失败")
//		fmt.Println(err)
//	}
//
//	fmt.Println("连接成功！！！")
//	user := User{Name: "as", Age: 43}
//	db.First(&user, "age = ?", "12") // 查找 code 字段值为 D42 的记录
//
//}
