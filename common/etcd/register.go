package etcd

import (
	"context"
	"fmt"
	"github.com/ricky-zhf/go-web/common/tools"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"time"
)

/*
服务注册
*/

const (
	ETCDPrefix = "Server_Register"
)

var (
	etcdRegister *EtcdRegister
)

type EtcdRegister struct {
	etcdClient *clientv3.Client
	leaseId    clientv3.LeaseID
	ctx        context.Context
	cancel     context.CancelFunc
	svrInfoMap map[string]map[string]string // service:[ip:port]
	mutex      sync.Mutex
}

// CreateLease 创建租约 expire 有效期/秒
func (e *EtcdRegister) CreateLease(expire int64) error {
	res, err := e.etcdClient.Grant(e.ctx, expire)
	if err != nil {
		log.Println("create etcd lease failed|error=", err)
		return err
	}
	log.Println("create etcd lease success...|res=", res)
	e.leaseId = res.ID
	return nil
}

// BindLease 绑定租约 将租约和对应的KEY-VALUE绑定
func (e *EtcdRegister) BindLease(key string, value string) error {
	res, err := e.etcdClient.Put(e.ctx, key, value, clientv3.WithLease(e.leaseId))
	if err != nil {
		log.Println("bind etcd lease failed|error=", err)
		return err
	}
	log.Println("bind etcd lease success...|res=", res)
	return nil
}

// KeepAlive 续租 发送心跳，表明服务正常
func (e *EtcdRegister) KeepAlive() (<-chan *clientv3.LeaseKeepAliveResponse, error) {
	resChan, err := e.etcdClient.KeepAlive(e.ctx, e.leaseId)
	if err != nil {
		log.Println("keepAlive failed,error=", resChan)
		return nil, err
	}

	return resChan, nil
}

// WatchLicense 监听续约
func (e *EtcdRegister) WatchLicense(eChan <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case l := <-eChan:
			// 续约成功这里会输出eChan
			log.Printf("watcher keepalive successfully|lience:%+v \n", l)
		case <-e.ctx.Done():
			_ = e.Close()
			log.Println("watcher keepalive end...")
			return
		}
	}
}

// Close 关闭EtcdRegister
func (e *EtcdRegister) Close() error {
	e.cancel()

	log.Printf("etcd closeing...EtcdRegister=%+v\n", e)
	// 撤销租约
	_, err := e.etcdClient.Revoke(e.ctx, e.leaseId)
	if err != nil {
		log.Println("etcd client revoke failed|err=", err)
	}

	return e.etcdClient.Close()
}

// RegisterServeice 注册服务 expire 过期时间
func (e *EtcdRegister) RegisterServeice(serviceName, ip, port, weight string, expire int64) {
	// 创建租约
	var err error
	if err = e.CreateLease(expire); err != nil {
		return
	}

	// 绑定租约
	key := generateKey(serviceName, ip, port)
	if err = e.BindLease(key, weight); err != nil {
		return
	}

	// 续租
	keepAliveChan, err := e.KeepAlive()
	if err != nil {
		return
	}

	// 监听续约
	go e.WatchLicense(keepAliveChan)
}

// RegisterAndDiscover 创建etcd register & discover
func RegisterAndDiscover(endpoints []string, expire int, serviceName, port, weight string, ttl int64) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Duration(expire) * time.Second,
	})
	if err != nil {
		log.Println("new etcd client failed,error=", err)
		return err
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	etcdRegister = &EtcdRegister{
		etcdClient: cli,
		ctx:        ctx,
		cancel:     cancelFunc,
		svrInfoMap: make(map[string]map[string]string, 0),
	}
	//defer etcdRegister.Close()

	// 服务注册
	etcdRegister.RegisterServeice(serviceName, tools.GetLocalIP(), port, weight, ttl)

	go etcdRegister.DiscoverService(serviceName)

	return nil
}

// ETCDKEY-blog_server-192.168.0.1-9090
func generateKey(serviceName, ip, port string) string {
	return fmt.Sprintf("%s:%s:%s:%s", ETCDPrefix, serviceName, ip, port)
}
