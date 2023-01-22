package etcd

import (
	"common/tools"
	"log"
	"testing"
)

func TestEtcdRegister_GetService(t *testing.T) {
	e, err := NewEtcdRegister([]string{"localhost:2379"}, 2)
	if err != nil {
		log.Println(err)
		return
	}
	defer e.Close()

	if err = e.RegisterServer("test_server123", tools.GetLocalIP(), "9091", "1", 2); err != nil {
		return
	}

	err = e.GetService(ETCDPrefix)
	log.Println("-----", err)
	tools.BlockMain()
}
