package Config

import (
	_ "github.com/go-sql-driver/mysql" // 必须要加，否则会找不到驱动
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func Conn() {
	// 为了完全支持utf-8编码，应该设置charset为utf8mb4
	// 为了能处理time.Time类型，应该包含参数parseTime
	db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("%#v", err)
		os.Exit(11)
	}
	DB = db
}

func Close()  {
	DB.Close()
}