package main

import (
	"context"
	account "github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
)

// AccountImpl implements the last service interface defined in the IDL.
type AccountImpl struct{}

// CreateAccount implements the AccountImpl interface.
func (s *AccountImpl) CreateAccount(ctx context.Context, req *account.CreateAccountRequest) (resp *account.CreateAccountResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteAccount implements the AccountImpl interface.
func (s *AccountImpl) DeleteAccount(ctx context.Context, req *account.DeleteAccountRequest) (resp *account.DeleteAccountResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateAccount implements the AccountImpl interface.
func (s *AccountImpl) UpdateAccount(ctx context.Context, req *account.UpdateAccountRequest) (resp *account.UpdateAccountResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetAccount implements the AccountImpl interface.
func (s *AccountImpl) MGetAccount(ctx context.Context, req *account.MGetAccountRequest) (resp *account.MGetAccountResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAccountByUserID implements the AccountImpl interface.
func (s *AccountImpl) GetAccountByUserID(ctx context.Context, req *account.GetAccountByUserIDRequest) (resp *account.GetAccountByUserIDResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateBill implements the AccountImpl interface.
func (s *AccountImpl) CreateBill(ctx context.Context, req *account.CreateBillRequest) (resp *account.CreateBillResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteBill implements the AccountImpl interface.
func (s *AccountImpl) DeleteBill(ctx context.Context, req *account.DeleteBillRequest) (resp *account.DeleteBillResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateBill implements the AccountImpl interface.
func (s *AccountImpl) UpdateBill(ctx context.Context, req *account.UpdateBillRequest) (resp *account.UpdateBillResponse, err error) {
	// TODO: Your code here...
	return
}

// MGetGillRequest implements the AccountImpl interface.
func (s *AccountImpl) MGetGillRequest(ctx context.Context, req *account.MGetBillRequest) (resp *account.MGetBillResponse, err error) {
	// TODO: Your code here...
	return
}

// GetBillByAccountID implements the AccountImpl interface.
func (s *AccountImpl) GetBillByAccountID(ctx context.Context, req *account.GetBillByAccountIDRequest) (resp *account.GetBillByAccountIDResponse, err error) {
	// TODO: Your code here...
	return
}
