package server

import (
	"context"
	"github.com/ricky-zhf/go-web/blog_server/dao"
	"github.com/ricky-zhf/go-web/common/pb"
	"log"
)

type BlogServer struct {
	pb.UnimplementedBlogServiceServer
}

func (s *BlogServer) GetBlog(ctx context.Context, in *pb.GetBlogRequest) (*pb.GetBlogResponse, error) {
	blogs, err := dao.GetBlogs()
	if err != nil {
		log.Println("failed to get blogs|err=", err)
		return nil, nil
	}
	var res []*pb.Blog
	defer log.Println("Get Res=", res, "err", err)
	for _, v := range blogs {
		res = append(res, &pb.Blog{
			Author:  v.Author,
			Title:   v.Title,
			Content: v.Content,
		})
	}

	return &pb.GetBlogResponse{Blog: res}, nil
}
