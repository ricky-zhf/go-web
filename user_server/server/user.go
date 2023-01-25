package server

import (
	"context"
	"github.com/ricky-zhf/go-web/common/pb/user"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (u *UserService) VerifyUserPwd(ctx context.Context, req *user.VerifyUserPwdRequest) (*user.VerifyUserPwdResponse, error) {
	res := user.ResOfPwd_Forbid
	if req.Name == "zhf" && req.Password == "123" {
		res = user.ResOfPwd_Pass
	}
	return &user.VerifyUserPwdResponse{ResOfPwd: res}, nil
}
