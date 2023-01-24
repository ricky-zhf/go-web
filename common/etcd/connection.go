package etcd

import (
	"fmt"
	"log"
	"math/rand"
)

func GetAddress(serviceName string) string {
	addrMap, ok := etcdTool.svrInfoMap[serviceName]
	if !ok {
		log.Println("GetAddress no info|serviceName=", serviceName)
		return ""
	}

	log.Println("GetAddress | addr=", addrMap)
	//随机取
	addrSlice := make([]string, 0)
	for k, v := range addrMap {
		addrSlice = append(addrSlice, fmt.Sprintf("%s:%s", k, v))
	}

	random := rand.Intn(len(addrSlice))
	resAddr := addrSlice[random]
	log.Printf("GetAddress end|addrSlice=%+v|random=%v|resAddr=%v", addrSlice, random, resAddr)
	return resAddr
}
