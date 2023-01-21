package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"user_server/config"
)

// DB 为了项目使用方便，有些变量可以定义为全局。
var (
	DB *gorm.DB
)

//链接mysql
func InitMySQL() error {
	common

	log.Println("init mysql start...")
	defer func() {
		log.Println("init mysql successfully...")
	}()

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.User, config.Conf.MySQL.Password, config.Conf.MySQL.Host, config.Conf.MySQL.Port, config.Conf.MySQL.DB)

	//"root:1234@tcp(mysqlv1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local" //mysqlv1
	DB, err := gorm.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
