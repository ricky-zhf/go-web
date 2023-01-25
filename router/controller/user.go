package controller

/*
1、json->rpc
2、获取etcd的Client
3、获取对应ip addr，然后请求
*/
//
//func HandlerUserRouter(r *gin.RouterGroup) {
//	r.POST("/GetUserAllBlogs", GetUserAllBlogs)
//}
//
//func GetUserAllBlogs(c *gin.Context) {
//	data, err := c.GetRawData()
//	if err != nil {
//		log.Println("GetUserAllBlogs GetRawData failed|err=", err)
//		return
//	}
//
//	userPb := &user.VerifyUserPwdRequest{}
//	if err = jsonpb.UnmarshalString(string(data), userPb); err != nil {
//		log.Println("GetUserAllBlogs UnmarshalString failed|err=", err)
//		return
//	}
//
//	address := etcd.GetAddress(config.Conf.Backends.UserService)
//	log.Println("GetUserAllBlogs GetAddress|addr=", address)
//
//	conn, err := rpc.GetRpcConn(address)
//	if err != nil {
//		log.Println("GetUserAllBlogs GetRpcConn failed|err=", err)
//		return
//	}
//	client := user.NewUserServiceClient(conn)
//
//	//使用服务定义好的方法
//	res, err := client.VerifyUserPwd(c, userPb)
//	if err != nil {
//		log.Println("GetBlog Failed|err=", err)
//		return
//	}
//	if res.ResOfPwd == user.ResOfPwd_Forbid{
//		log.Println("pwd forbid")
//		c.JSON(http.StatusOK, gin.H{
//			"code": 4040,
//			"msg":  "pwd forbid",
//		})
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": 200,
//		"msg":  "success",
//	})
//}
