package main

import (
	"github.com/ricky-zhf/go-web/common/etcd"
	"github.com/ricky-zhf/go-web/router/config"
	"github.com/ricky-zhf/go-web/router/router"
	"log"
)

func main() {

	var err error
	if err = config.InitConfig(); err != nil {
		log.Fatalln("init config failed|err=", err)
	}

	router.SetupRouter()

	if err = etcd.RegisterAndDiscover(
		config.Conf.Etcd.Endpoints, 5, config.Conf.Service.Name,
		config.Conf.Service.Port, config.Conf.Service.Weight, 5,
	); err != nil {
		log.Fatalln("init etcd failed|err=", err)
	}
}
