package pack

import (
	"errors"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/dal/db"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

func Account(model *db.Account) *account.Account {
	return &account.Account{
		AccountId:   int64(model.ID),
		AccountName: model.AccountName,
		Balance:     model.Balance,
	}
}

func Accounts(model []*db.Account) []*account.Account {
	res := make([]*account.Account, 0, len(model))
	for _, item := range model {
		res = append(res, Account(item))
	}
	return res
}

func BuildBaseResp(err error) *account.BaseResp {
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

func baseResp(err errno.ErrNo) *account.BaseResp {
	return &account.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
	}
}
