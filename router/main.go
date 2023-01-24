package main

import (
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
}
