package main

import (
	"blog_server/dao"
	"blog_server/pb"
	. "blog_server/server"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		log.Panicln("failed to init mysql|err=", err)
	}
	log.Println("InitMySQL successful")
	defer dao.CloseDB()

	listen, err := net.Listen("tcp", ":9090") //这里不能使用127.0.0.1,因为要放到docker里会导致只能在docker里面才能访问服务
	if err != nil {
		log.Panicln("net.Listen failed|err=", err)
	}
	log.Println("Listen 9090 successful")
	defer listen.Close()

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterBlogServiceServer(server, &BlogServer{})

	log.Println("start listening...")

	//startRoute()

	err = server.Serve(listen)
	if err != nil {
		log.Panicln("failed to server|err=", err)
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
