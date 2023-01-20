package main

import (
	. "bubble/blog_server/server"
	"bubble/pb"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("todo server start")
	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		fmt.Printf("net.Listen failed|err=%v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterBlogServer(server, &BlogServer{})
	err = server.Serve(listen)
	if err != nil {
		fmt.Printf("failed to server|err=%v\n", err)
	}
	fmt.Println("todo server started successfully")

}
