package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ricky-zhf/go-web/blog_server/config"
	"log"
)

// DB 为了项目使用方便，有些变量可以定义为全局。
var (
	DB *gorm.DB
)

// InitMySQL 链接mysql
func InitMySQL() error {
	log.Println("init mysql start...")
	defer func() {
		log.Println("init mysql successfully...")
	}()

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.User, config.Conf.MySQL.Password, config.Conf.MySQL.Host, config.Conf.MySQL.Port, config.Conf.MySQL.DB)

	var err error
	if DB, err = gorm.Open("mysql", dbConn); err != nil {
		return err
	}

	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
