package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/ricky-zhf/go-web/common/etcd"
	rpc "github.com/ricky-zhf/go-web/common/grpc"
	"github.com/ricky-zhf/go-web/common/pb/blog"
	"github.com/ricky-zhf/go-web/common/pb/user"
	"github.com/ricky-zhf/go-web/router/config"
	"log"
	"net/http"
)

func HandlerBlogRouter(r *gin.RouterGroup) {
	r.POST("/GetUserAllBlogs", GetUserAllBlogs)
}

func GetUserAllBlogs(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Println("GetUserAllBlogs GetRawData failed|err=", err)
		return
	}

	userPb := &user.VerifyUserPwdRequest{}
	if err = jsonpb.UnmarshalString(string(data), userPb); err != nil {
		log.Println("GetUserAllBlogs UnmarshalString failed|err=", err)
		return
	}

	//这里可以直接集成到common包里，为了清晰流程放在业务代码中
	address := etcd.GetAddress(config.Conf.Backends.UserService)
	log.Println("GetUserAllBlogs GetAddress|addr=", address)

	conn, err := rpc.GetRpcConn(address)
	if err != nil {
		log.Println("GetUserAllBlogs GetRpcConn failed|err=", err)
		return
	}
	client := user.NewUserServiceClient(conn)

	res, err := client.VerifyUserPwd(c, userPb)
	if err != nil {
		log.Println("GetBlog Failed|err=", err)
		return
	}
	if res.ResOfPwd == user.ResOfPwd_Forbid {
		log.Println("pwd forbid")
		c.JSON(http.StatusOK, gin.H{
			"code": 4040,
			"msg":  "pwd forbid",
		})
		return
	}

	//通过user和password校验，获取blogs
	conn, err = rpc.GetRpcConn(etcd.GetAddress(config.Conf.Backends.BlogService))
	if err != nil {
		log.Println("GetUserAllBlogs GetRpcConn failed|err=", err)
		return
	}
	req := blogGo.GetUserAllBlogsRequest{UserName: userPb.Name}
	blogs, err := blogGo.NewBlogServiceClient(conn).GetUserAllBlogs(c, &req)
	if err != nil {
		log.Println("GetUserAllBlogs NewBlogServiceClient failed|err=", err)
		return
	}
	data, _ = json.Marshal(blogs)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": string(data),
	})
}
