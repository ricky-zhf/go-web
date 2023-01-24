package server

import (
	"context"
	"github.com/ricky-zhf/go-web/common/pb/blog"
	"github.com/ricky-zhf/go-web/common/pb/user"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u *UserService) GetAllUserBlogs(ctx context.Context, in *user.GetAllUserBlogsRequest) (*user.GetAllUserBlogsResponse, error) {
	blogs := []*blogGo.Blog{
		{
			Author:  "zhf",
			Title:   "Title Get ALl Blogs",
			Content: "Content Get All Blogs",
		},
	}
	return &user.GetAllUserBlogsResponse{
		Blog: blogs,
	}, nil
}
