package etcd

import (
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strings"
)

/*
服务发现:
	- GetService获取etcd中前缀为serviceName的所有key,value
	- 随后逐个updateSvrMap到EtcdRegister中的svrInfoMap中
	- 随后watchService方法监控etcd中前缀为serviceName是否变动，并更新到svrInfoMap中
*/

func (e *EtcdRegister) DiscoverService(serviceName string) {
	serviceName = ETCDPrefix + serviceName
	// Client.KV 是一个 interface ，提供了关于 K-V 操作的所有方法.
	kv := clientv3.NewKV(e.etcdClient)
	resp, err := kv.Get(e.ctx, serviceName, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("get kv from etcd failed|serName=%v|err=%v\n", serviceName, err)
	}

	// update svrMap
	for _, kv := range resp.Kvs {
		e.updateSvrMap(string(kv.Key))
	}

	// watch service change
	go e.watchService()
}

func (e *EtcdRegister) watchService() {
	watchChan := e.etcdClient.Watch(e.ctx, ETCDPrefix, clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT: //PUT事件，目录下有了新key
				log.Println("watch service key changed|key=", string(event.Kv.Key))
				e.updateSvrMap(string(event.Kv.Key))
			case mvccpb.DELETE: //DELETE事件，目录中有key被删掉(Lease过期，key 也会被删掉)
				log.Println("watch service key deleted|key=", string(event.Kv.Key))
				e.mutex.Lock()
				sli := splitKey(string(event.Kv.Key))
				delete(e.svrInfoMap, sli[1])
				e.mutex.Unlock()
			}
			log.Println("map=", e.svrInfoMap)
		}
	}
}

func (e *EtcdRegister) updateSvrMap(s string) {
	//ETCDKEY-blog_server-192.168.0.1-9090
	sli := splitKey(s)
	if len(sli) != 4 {
		log.Printf("something wrong in resolving kvs|key=%v", s)
		return
	}

	e.mutex.Lock()
	serName := sli[1]
	serIp := sli[2]
	serPort := sli[3]
	if _, ok := e.svrInfoMap[serName]; ok {
		e.svrInfoMap[serName][serIp] = serPort
	} else {
		m := map[string]string{
			serIp: serPort,
		}
		e.svrInfoMap[serName] = m
	}
	e.mutex.Unlock()

}

func splitKey(key string) []string {
	return strings.Split(key, "-")
}
