package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ricky-zhf/go-web/blog_server/config"
	"github.com/ricky-zhf/go-web/blog_server/dao"
	se "github.com/ricky-zhf/go-web/blog_server/server"
	"github.com/ricky-zhf/go-web/common/etcd"
	"github.com/ricky-zhf/go-web/common/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	var err error
	if err = config.InitConfig(); err != nil {
		log.Fatalln("init config failed|err=", err)
	}

	if err = dao.InitMySQL(); err != nil {
		log.Fatalln("init mysql failed|err=", err)
	}
	defer dao.CloseDB()

	if err = etcd.RegisterAndDiscover(
		config.Conf.Etcd.Endpoints, 5, config.Conf.Service.Name,
		config.Conf.Service.Port, config.Conf.Service.Weight, 5,
	); err != nil {
		log.Fatalln("init etcd failed|err=", err)
	}

	//这里不能使用127.0.0.1,因为要放到docker里会导致只能在docker里面才能访问服务
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Conf.Service.Port))
	if err != nil {
		log.Fatalln("net.Listen failed|err=", err)
	}
	defer listen.Close()

	server := grpc.NewServer()
	reflection.Register(server) //供grpcurl命令使用
	pb.RegisterBlogServiceServer(server, &se.BlogServer{})

	log.Println("grpc register end...")

	if err = server.Serve(listen); err != nil {
		log.Fatalln("failed to server|err=", err)
	}
}

func startRoute() {
	// 创建路由
	r := gin.Default()
	// 绑定路由规则，执行的函数（gin.Context，封装了request和response）
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080，Run("里面不指定端口号默认为8080")
	err := r.Run(":9090")

	log.Println("start route", err)
}
