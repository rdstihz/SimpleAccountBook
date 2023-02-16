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

type BillParam struct {
	UserID    int64   `json:"user_id"`
	AccountID int64   `json:"account_id"`
	Type      int64   `json:"type"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	Comment   string  `json:"comment"`
}

func CreateBillHandler(ctx context.Context, c *app.RequestContext) {
	//从jwt token payload中获取用户id
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constants.IdentityKey].(float64))
	//获取account id
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//获取传入的bill的信息
	var billValues BillParam
	err = c.Bind(&billValues)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//调用rpc创建帐单
	err = rpc.CreateBill(&account.CreateBillRequest{Bill: &account.Bill{
		UserId:    userID,
		AccountId: accountID,
		Type:      billValues.Type,
		Amount:    billValues.Amount,
		Category:  billValues.Category,
		Comment:   billValues.Comment,
	}})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func DeleteBillHandler(ctx context.Context, c *app.RequestContext) {
	//bill id
	billID, err := strconv.ParseInt(c.Param("bill_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//调用rpc删除帐单
	err = rpc.DeleteBill(&account.DeleteBillRequest{BillId: billID})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func UpdateBillHandler(ctx context.Context, c *app.RequestContext) {
	//获取bill id
	billID, err := strconv.ParseInt(c.Param("bill_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//获取传入的bill的信息
	var billValues BillParam
	err = c.Bind(&billValues)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//调用rpc修改帐单
	err = rpc.UpdateBill(&account.UpdateBillRequest{Bill: &account.Bill{
		BillId:    billID,
		UserId:    billValues.UserID,
		AccountId: billValues.AccountID,
		Type:      billValues.Type,
		Amount:    billValues.Amount,
		Category:  billValues.Category,
		Comment:   billValues.Comment,
	}})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func GetBillInfoHandler(ctx context.Context, c *app.RequestContext) {
	//获取bill id
	billID, err := strconv.ParseInt(c.Param("bill_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//调用rpc
	bills, err := rpc.MGetBill(&account.MGetBillRequest{BillIds: []int64{billID}})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	if len(bills) == 0 {
		SendResponse(c, errno.BillNotFoundErr, nil)
		return
	}
	SendResponse(c, errno.Success, bills[0])
}

func ListBillHandler(ctx context.Context, c *app.RequestContext) {
	//获取account id
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	//调用rpc
	bills, err := rpc.GetBillByAccountID(&account.GetBillByAccountIDRequest{AccountId: accountID})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, bills)
}
