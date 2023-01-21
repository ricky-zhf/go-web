package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"time"
	"user_server/config"
	"user_server/dao"
	"user_server/pb"
)

func main() {
	var err error
	if err = config.InitConfig(); err != nil {
		log.Println("init config failed|err=", err)
	}

	if err = InitEtcd(); err != nil {
		log.Println("init etcd failed|err=", err)
	}

	if err = dao.InitMySQL(); err != nil {
		log.Println("failed to init mysql|err=", err)
	}

	//先拨号, withblock， 拨号成功才往下运行
	conn, err := grpc.Dial("blog_server:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("client cannot dial grpc server|err=", err)
	}
	defer conn.Close()

	getAddr()

	//新建客户端，将grpc链接放入，然后返回服务中的client接口。
	client := pb.NewBlogServiceClient(conn)

	//client就可以使用服务定义好的方法
	blog, err := client.GetBlog(context.Background(), &pb.GetBlogRequest{Title: "123"})
	if err != nil {
		log.Println("GetBlog Failed|err=", err)
	}

	log.Println("blog:", blog)

	log.Println("start user server successfully...")

	block()
}

func InitEtcd() error {
	var err error
	log.Println("start init etcd...")
	defer func() {
		log.Println("end init etcd...", err)
	}()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   config.Conf.Etcd.Endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("connect to etcd failed| err=%v\n", err)
		return err
	}
	log.Println("etcd connect successfully...")
	defer cli.Close()
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "test...", "testSuc...")
	cancel()
	if err != nil {
		log.Printf("put to etcd failed| err=%v\n", err)
		return err
	}
	return nil
}

func getAddr() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("get addr failed|err=", err)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				log.Println("local ip addr=", ipnet.IP.String())
			}
		}
	}
}

func block() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		log.Println("listen failed, err:", err)
		return
	}
	//如果监听成功，则无限循环，监听是否有链接
	for true {
		_, err = listen.Accept() //接受listen所监听的端口传入的链接
		if err != nil {
			log.Println("accept failed , err:", err)
			continue
		}
		//没有错误，链接建立完成，处理连接的请求。
	}
}
