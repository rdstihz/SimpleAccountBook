package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	userapi "github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/user/userservice"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

var userClient userservice.Client

// 初始化user rpc连接
func initUserRPC() {
	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithHostPorts("127.0.0.1:18881"),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

// CreateUser 调用rpc创建用户
func CreateUser(req *userapi.CreateUserRequest) error {
	resp, err := userClient.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return nil
}

// CheckUser 调用rpc检查用户名和密码是否正确，如果验证成功，返回用户id
func CheckUser(req *userapi.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(context.Background(), req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.User.UserId, nil
}

// MGetUser 调用rpc，根据user id查询用户信息
func MGetUser(req *userapi.MGetUserRequest) ([]*userapi.User, error) {
	resp, err := userClient.MGetUser(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Users, nil
}
