package main

import (
	"common/etcd"
	"common/tools"
	"log"
)

type C struct {
	SS string
}

func main() {
	e, err := etcd.NewEtcdRegister([]string{"localhost:2379"}, 2)
	if err != nil {
		log.Println(err)
		return
	}
	defer e.Close()

	if err = e.RegisterServer("test-server", tools.GetLocalIP(), "9091", "1", 2); err != nil {
		log.Println("fffferr=", err)
		return
	}

	tools.BlockMain()
}
