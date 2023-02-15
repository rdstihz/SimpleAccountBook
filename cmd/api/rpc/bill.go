package rpc

import (
	"context"
	accountapi "github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

// CreateBill 调用rpc, 创建帐单
func CreateBill(req *accountapi.CreateBillRequest) error {
	resp, err := accountClient.CreateBill(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// DeleteBill 调用rpc，删除帐单
func DeleteBill(req *accountapi.DeleteBillRequest) error {
	resp, err := accountClient.DeleteBill(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// UpdateBill 调用rpc，修改帐单
func UpdateBill(req *accountapi.UpdateBillRequest) error {
	resp, err := accountClient.UpdateBill(context.Background(), req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// MGetBill 调用rpc，根据id查询帐单
func MGetBill(req *accountapi.MGetBillRequest) ([]*accountapi.Bill, error) {
	resp, err := accountClient.MGetBill(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Bills, nil
}

// GetBillByAccountID 根据account id，查询account下的所有帐单
func GetBillByAccountID(req *accountapi.GetBillByAccountIDRequest) ([]*accountapi.Bill, error) {
	resp, err := accountClient.GetBillByAccountID(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Bills, nil
}
