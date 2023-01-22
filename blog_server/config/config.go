package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type MySQL struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type Config struct {
	MySQL   MySQL
	Etcd    Etcd
	Service Service
}

type Service struct {
	Name   string
	Port   string
	Weight string
}

type Etcd struct {
	Endpoints []string
}

var Conf Config

func InitConfig() (err error) {
	log.Println("start init config...")
	defer func() {
		log.Println("end init etcd...")
	}()

	//todo delete
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("读取应用目录失败|err=", err)
	}
	log.Println("work dir=", workDir)

	viper.SetConfigFile("config/config.yaml")

	// 读取配置信息
	if err = viper.ReadInConfig(); err != nil {
		log.Println("read config failed|err=", err)
	}
	if err = viper.Unmarshal(&Conf); err != nil {
		log.Println("unmarshal config failed|err=", err)
	}

	allSets := viper.AllSettings()
	log.Println("AllSettings=", allSets)

	return nil
}
