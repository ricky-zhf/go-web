package main

import (
	"bubble/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//先拨号
	//withblock， 拨号成功才往下运行
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("client cannot dial grpc sercer")
	}
	defer conn.Close()
	//新建客户端，将grpc链接放入，然后返回服务中的client接口。
	client := pb.NewBlogClient(conn)
	//client就可以使用服务定义好的方法
	blog, err := client.GetBlog(context.Background(), &pb.GetBlogRequest{Title: "123"})
	if err != nil {
		fmt.Printf("GetBlog Failed|err=%v\n", err)
		return
	}
	fmt.Println("blog:", blog)
}
