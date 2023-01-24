package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"user_server/config"
)

var (
	DB *gorm.DB
)

func InitMySQL() error {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.MySQL.User, config.Conf.MySQL.Password, config.Conf.MySQL.Host,
		config.Conf.MySQL.Port, config.Conf.MySQL.DB)

	var err error
	if DB, err = gorm.Open("mysql", dbConn); err != nil {
		return err
	}
	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
