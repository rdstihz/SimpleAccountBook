package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/rdstihz/SimpleAccountBook/cmd/api/rpc"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

type UserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterHandler 用户注册handler
func RegisterHandler(ctx context.Context, c *app.RequestContext) {
	//获取传入的用户名密码
	var registerValues UserParam
	err := c.Bind(&registerValues)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	username := registerValues.Username
	password := registerValues.Password
	if len(username) == 0 || len(password) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	//调用rpc创建用户
	err = rpc.CreateUser(&user.CreateUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// GetUserInfoHandler 查询用户信息handler
func GetUserInfoHandler(ctx context.Context, c *app.RequestContext) {
	//从jwt token payload中获取用户id
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	//调用rpc查询用户信息
	users, err := rpc.MGetUser(&user.MGetUserRequest{UserIds: []int64{userID}})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	if len(users) == 0 {
		SendResponse(c, errno.UserNotFoundErr, nil)
	}
	SendResponse(c, errno.Success, users[0])
}
