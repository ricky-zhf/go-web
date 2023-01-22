package etcd

import (
	"github.com/ricky-zhf/go-web/common/tools"
	"log"
	"testing"
)

func TestEtcdRegister_GetService(t *testing.T) {
	err := RegisterAndDiscover([]string{"localhost:2379"}, 2, "", "2379", "1", 2)
	if err != nil {
		log.Println(err)
		return
	}
	tools.BlockMain()
}
