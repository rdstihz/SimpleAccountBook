package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/rdstihz/SimpleAccountBook/cmd/user/dal/db"
	"github.com/rdstihz/SimpleAccountBook/cmd/user/pack"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
	"io"
)

// CreateUser 创建用户
func CreateUser(ctx context.Context, username, password string) error {
	// 查询用户名是否已被使用
	users, err := db.QueryUser(ctx, username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	//计算密码md5
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return err
	}
	passwordMd5 := fmt.Sprintf("%x", h.Sum(nil))
	//写入数据库
	return db.CreateUser(ctx, []*db.User{
		&db.User{
			Username: username,
			Password: passwordMd5,
		},
	})
}

// CheckUser 验证用户
func CheckUser(ctx context.Context, username, password string) (*user.User, error) {
	// 在数据库中查找用户
	users, err := db.QueryUser(ctx, username)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.AuthorizationFailedErr
	}

	//验证密码是否正确
	h := md5.New()
	if _, err := io.WriteString(h, password); err != nil {
		return nil, err
	}
	passwordMd5 := fmt.Sprintf("%x", h.Sum(nil))
	if passwordMd5 != users[0].Password {
		return nil, errno.AuthorizationFailedErr
	}
	return pack.GormUserToUser(users[0]), nil
}

// MGetUser 查询用户
func MGetUser(ctx context.Context, ids []int64) ([]*user.User, error) {
	users, err := db.MGetUsers(ctx, ids)
	if err != nil {
		return nil, err
	}
	return pack.GormUsersToUsers(users), nil
}
