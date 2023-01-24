package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/ricky-zhf/go-web/common/etcd"
	"github.com/ricky-zhf/go-web/common/pb/user"
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

	address := etcd.GetAddress("Server_Register-blog.UserService")
	log.Println("GetUserAllBlogs GetAddress|addr=", address)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
