package main

import (
	"context"
	"github.com/ricky-zhf/go-web/common/etcd"
	"github.com/ricky-zhf/go-web/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"user_server/config"
	"user_server/dao"
)

func main() {
	var err error
	if err = config.InitConfig(); err != nil {
		log.Fatalln("init config failed|err=", err)
	}

	if err = etcd.RegisterAndDiscover(
		config.Conf.Etcd.Endpoints, 5, config.Conf.Service.Name,
		config.Conf.Service.Port, config.Conf.Service.Weight, 5,
	); err != nil {
		log.Fatalln("init etcd failed|err=", err)
	}

	if err = dao.InitMySQL(); err != nil {
		log.Fatalln("failed to init mysql|err=", err)
	}
	defer dao.CloseDB()

	// todo 改造
	conn, err := grpc.Dial("blog_server:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("client cannot dial grpc server|err=", err)
	}
	defer conn.Close()

	//新建客户端，将grpc链接放入，然后返回服务中的client接口。
	client := pb.NewBlogServiceClient(conn)

	//使用服务定义好的方法
	res, err := client.GetBlog(context.Background(), &pb.GetBlogRequest{Title: "123"})
	if err != nil {
		log.Println("GetBlog Failed|err=", err)
	}
	log.Println("Get res=", res)
}
