package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"github.com/rdstihz/SimpleAccountBook/cmd/api/rpc"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
	"strconv"
)

type AccountParam struct {
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func CreateAccountHandler(ctx context.Context, c *app.RequestContext) {
	//从jwt token payload中获取用户id
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))

	//获取传入的account_name和balance
	var accountValues AccountParam
	err := c.Bind(&accountValues)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	accountName := accountValues.AccountName
	balance := accountValues.Balance
	// 调用rpc创建account
	err = rpc.CreateAccount(&account.CreateAccountRequest{
		UserId:      userID,
		AccountName: accountName,
		Balance:     balance,
	})

	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func DeleteAccountHandler(ctx context.Context, c *app.RequestContext) {
	//从jwt token payload中获取用户id
	//claims := jwt.ExtractClaims(ctx, c)
	//userID := int64(claims[constants.IdentityKey].(float64))

	//获取要删除的account id
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	// 调用rpc创建account
	err = rpc.DeleteAccount(&account.DeleteAccountRequest{AccountId: accountID})

	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func UpdateAccountHandler(ctx context.Context, c *app.RequestContext) {
	//从jwt token payload中获取用户id
	//claims := jwt.ExtractClaims(ctx, c)
	//userID := int64(claims[constants.IdentityKey].(float64))
	//获取要修改的account id
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//获取传入的account_name和balance
	var accountValues AccountParam
	err = c.Bind(&accountValues)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	accountName := accountValues.AccountName
	balance := accountValues.Balance
	// 调用rpc创建account
	err = rpc.UpdateAccount(&account.UpdateAccountRequest{
		AccountId:   accountID,
		AccountName: accountName,
		Balance:     balance,
	})

	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetAccountInfoHandler(ctx context.Context, c *app.RequestContext) {
	//获取要查询的account id
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//调用rpc查询account信息
	accounts, err := rpc.MGetAccount(&account.MGetAccountRequest{AccountIds: []int64{accountID}})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	if len(accounts) == 0 {
		SendResponse(c, errno.AccountNotFoundErr, nil)
		return
	}
	SendResponse(c, errno.Success, accounts[0])
}

func ListAccountHandler(ctx context.Context, c *app.RequestContext) {
	//从jwt token payload中获取用户id
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	// 调用rpc查询用户所有account
	accounts, err := rpc.GetAccountByUserID(&account.GetAccountByUserIDRequest{UserId: userID})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, accounts)
}
