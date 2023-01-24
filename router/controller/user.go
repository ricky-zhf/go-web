package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/ricky-zhf/go-web/common/etcd"
	"github.com/ricky-zhf/go-web/common/pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

const (
	GetUserAllBlogsUri = "/GetUserAllBlogs"
)

/*
1、json->rpc
2、获取etcd的Client
3、获取对应ip addr，然后请求
*/

func HandlerUserRouter(r *gin.RouterGroup) {
	r.POST("/GetUserAllBlogs", GetUserAllBlogs)
}

func GetUserAllBlogs(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Println("GetUserAllBlogs GetRawData failed|err=", err)
		return
	}

	userPb := &user.GetAllUserBlogsRequest{}
	if err = jsonpb.UnmarshalString(string(data), userPb); err != nil {
		log.Println("GetUserAllBlogs UnmarshalString failed|err=", err)
		return
	}

	address := etcd.GetAddress("blog.UserService")
	log.Println("GetUserAllBlogs GetAddress|addr=", address)

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("client cannot dial grpc server|err=", err)
	}
	client := user.NewUserServiceClient(conn)

	//使用服务定义好的方法
	res, err := client.GetAllUserBlogs(c, userPb)
	if err != nil {
		log.Println("GetBlog Failed|err=", err)
	}
	log.Println("ressssss", res)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
