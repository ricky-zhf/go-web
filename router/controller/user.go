package controller

import "github.com/gin-gonic/gin"

/*
1、json->rpc
2、获取etcd的Client
3、获取对应ip addr，然后请求
*/

func HandlerUserRouter(r *gin.RouterGroup) {
	r.GET("/GetUserAllBlogs", GetUserAllBlogs)
}

func GetUserAllBlogs(r *gin.Context) {

}
