package rpc

import (
	"context"
	accountapi "github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/account/accountservice"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

var accountClient accountservice.Client

// 初始化account rpc 连接
func initAccountRPC() {
	c, err := accountservice.NewClient(constants.AccountServiceName)
	if err != nil {
		panic(err)
	}
	accountClient = c
}

// CreateAccount 调用rpc,创建account
func CreateAccount(req *accountapi.CreateAccountRequest) error {
	resp, err := accountClient.CreateAccount(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// DeleteAccount 调用rpc，删除account
func DeleteAccount(req *accountapi.DeleteAccountRequest) error {
	resp, err := accountClient.DeleteAccount(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// UpdateAccount 调用rpc,修改account
func UpdateAccount(req *accountapi.UpdateAccountRequest) error {
	resp, err := accountClient.UpdateAccount(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// MGetAccount 调用rpc，查询指定id的account信息
func MGetAccount(req *accountapi.MGetAccountRequest) ([]*accountapi.Account, error) {
	resp, err := accountClient.MGetAccount(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Accounts, nil
}

// GetAccountByUserID 调用rpc，根据user id，查询用户的所有account
func GetAccountByUserID(req *accountapi.GetAccountByUserIDRequest) ([]*accountapi.Account, error) {
	resp, err := accountClient.GetAccountByUserID(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Accounts, nil
}
