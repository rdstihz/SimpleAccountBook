package service

import (
	"context"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/dal/db"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/pack"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
)

// CreateAccount 创建帐户
func CreateAccount(ctx context.Context, userID int64, accountName string, balance float64) error {
	return db.CreateAccount(ctx, []*db.Account{
		&db.Account{
			UserId:      userID,
			AccountName: accountName,
			Balance:     balance,
		},
	})
}

// DeleteAccount 删除帐户
func DeleteAccount(ctx context.Context, accountID int64) error {
	return db.DeleteAccount(ctx, accountID)
}

// UpdateAccount 修改帐户
func UpdateAccount(ctx context.Context, accountID int64, accountName string, balance float64) error {
	return db.UpdateAccount(ctx, accountID, &db.Account{
		AccountName: accountName,
		Balance:     balance,
	})
}

// MGetAccount 查询帐户
func MGetAccount(ctx context.Context, accountIDs []int64) ([]*account.Account, error) {
	models, err := db.MGetAccount(ctx, accountIDs)
	if err != nil {
		return nil, err
	}
	return pack.Accounts(models), nil
}

// GetAccountByUserID 查询用户的所有帐户
func GetAccountByUserID(ctx context.Context, userID int64) ([]*account.Account, error) {
	models, err := db.GetAccountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return pack.Accounts(models), nil
}
