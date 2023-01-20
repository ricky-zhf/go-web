package server

import (
	pb "bubble/pb"
	"context"
)

type BlogServer struct {
	pb.UnimplementedBlogServer
}

func (s *BlogServer) GetBlog(ctx context.Context, in *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	return &pb.GetBlogResponse{
		Author:  "zhf",
		Title:   "title",
		Content: "zhf's blog",
	}, nil
}
