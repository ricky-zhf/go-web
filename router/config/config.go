package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type BackendServices struct {
	UserService string
	BlogService string
}

type Config struct {
	UserServer string
	BlogServer string
	Etcd       Etcd
	Service    Service
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
		log.Println("failed to get wd|err=", err)
	}
	log.Println("work dir=", workDir)

	// 通过docker部署后，workDir是/ ，config目录已经通过dockerFile复制到/下面了，所以直接用config路径即可。
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
