package main

import (
	"context"
	"github.com/rdstihz/SimpleAccountBook/cmd/user/pack"
	"github.com/rdstihz/SimpleAccountBook/cmd/user/service"
	user "github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)
	username := req.Username
	password := req.Password
	if len(username) == 0 || len(password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.CreateUser(ctx, username, password)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	resp = new(user.MGetUserResponse)
	ids := req.UserIds
	if len(ids) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	users, err := service.MGetUser(ctx, ids)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)
	username := req.Username
	password := req.Password
	if len(username) == 0 || len(password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	user, err := service.CheckUser(ctx, username, password)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.User = user
	return
}
