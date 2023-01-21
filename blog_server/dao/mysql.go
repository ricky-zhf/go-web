package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB 为了项目使用方便，有些变量可以定义为全局。
var (
	DB *gorm.DB
)

//链接mysql
func InitMySQL() (err error) {
	dsn := "root:1234@tcp(mysqlv1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local" //mysqlv1
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
