package main

import (
	"context"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/pack"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/service"
	account "github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

// AccountImpl implements the last service interface defined in the IDL.
type AccountImpl struct{}

// CreateAccount implements the AccountImpl interface.
func (s *AccountImpl) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (resp *account.CreateAccountResponse, err error) {
	resp = new(account.CreateAccountResponse)

	userID := req.UserId
	accountName := req.AccountName
	balance := req.Balance

	if len(accountName) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.CreateAccount(ctx, userID, accountName, balance)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// DeleteAccount implements the AccountImpl interface.
func (s *AccountImpl) DeleteAccount(ctx context.Context, req *account.DeleteAccountRequest) (resp *account.DeleteAccountResponse, err error) {
	resp = new(account.DeleteAccountResponse)
	accountID := req.AccountId
	if accountID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	err = service.DeleteAccount(ctx, accountID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// UpdateAccount implements the AccountImpl interface.
func (s *AccountImpl) UpdateAccount(ctx context.Context, req *account.UpdateAccountRequest) (resp *account.UpdateAccountResponse, err error) {
	resp = new(account.UpdateAccountResponse)
	accountID := req.AccountId
	accountName := req.AccountName
	balance := req.Balance

	if accountID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.UpdateAccount(ctx, accountID, accountName, balance)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetAccount implements the AccountImpl interface.
func (s *AccountImpl) MGetAccount(ctx context.Context, req *account.MGetAccountRequest) (resp *account.MGetAccountResponse, err error) {
	resp = new(account.MGetAccountResponse)

	accountIDs := req.AccountIds
	if len(accountIDs) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	accounts, err := service.MGetAccount(ctx, accountIDs)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Accounts = accounts
	return
}

// GetAccountByUserID implements the AccountImpl interface.
func (s *AccountImpl) GetAccountByUserID(ctx context.Context, req *account.GetAccountByUserIDRequest) (resp *account.GetAccountByUserIDResponse, err error) {
	resp = new(account.GetAccountByUserIDResponse)

	userID := req.UserId
	if userID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	accounts, err := service.GetAccountByUserID(ctx, userID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Accounts = accounts
	return
}

// CreateBill implements the AccountImpl interface.
func (s *AccountImpl) CreateBill(ctx context.Context, req *account.CreateBillRequest) (resp *account.CreateBillResponse, err error) {
	resp = new(account.CreateBillResponse)

	bill := req.Bill
	if bill.BillId != 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}

	err = service.CreateBill(ctx, bill)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// DeleteBill implements the AccountImpl interface.
func (s *AccountImpl) DeleteBill(ctx context.Context, req *account.DeleteBillRequest) (resp *account.DeleteBillResponse, err error) {
	resp = new(account.DeleteBillResponse)
	billID := req.BillId
	if billID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	err = service.DeleteBill(ctx, billID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// UpdateBill implements the AccountImpl interface.
func (s *AccountImpl) UpdateBill(ctx context.Context, req *account.UpdateBillRequest) (resp *account.UpdateBillResponse, err error) {
	resp = new(account.UpdateBillResponse)

	bill := req.Bill
	if bill.BillId <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	err = service.UpdateBill(ctx, bill)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// MGetBill implements the AccountImpl interface.
func (s *AccountImpl) MGetBill(ctx context.Context, req *account.MGetBillRequest) (resp *account.MGetBillResponse, err error) {
	resp = new(account.MGetBillResponse)
	billIDs := req.BillIds
	if len(billIDs) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	bills, err := service.MGetBill(ctx, billIDs)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Bills = bills
	return
}

// GetBillByAccountID implements the AccountImpl interface.
func (s *AccountImpl) GetBillByAccountID(ctx context.Context, req *account.GetBillByAccountIDRequest) (resp *account.GetBillByAccountIDResponse, err error) {
	resp = new(account.GetBillByAccountIDResponse)
	accountID := req.AccountId
	if accountID <= 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return
	}
	bills, err := service.GetBillByAccountID(ctx, accountID)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Bills = bills
	return
}
