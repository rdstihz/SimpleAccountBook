package pack

import (
	"errors"
	"github.com/rdstihz/SimpleAccountBook/cmd/user/dal/db"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/user"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

func GormUserToUser(u *db.User) *user.User {
	return &user.User{
		UserId:   int64(u.ID),
		Username: u.Username,
	}
}

func GormUsersToUsers(u []*db.User) []*user.User {
	res := make([]*user.User, 0, len(u))
	for _, item := range u {
		res = append(res, GormUserToUser(item))
	}
	return res
}

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
