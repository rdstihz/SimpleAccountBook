package db

import (
	"context"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"gorm.io/gorm"
	"log"
)

type Account struct {
	gorm.Model
	UserId      int64   `json:"user_id"`
	AccountName string  `json:"account_name"`
	Balance     float64 `json:"balance"`
}

func (a *Account) TableName() string {
	return constants.AccountTableName
}

func CreateAccount(ctx context.Context, accounts []*Account) error {
	return DB.WithContext(ctx).Create(accounts).Error
}

func DeleteAccount(ctx context.Context, accountID int64) error {
	return DB.WithContext(ctx).Delete(&Account{}, accountID).Error
}

func UpdateAccount(ctx context.Context, accountID int64, account *Account) error {
	return DB.WithContext(ctx).Model(&Account{Model: gorm.Model{ID: uint(accountID)}}).Updates(account).Error
}

func MGetAccount(ctx context.Context, accountIDs []int64) ([]*Account, error) {
	res := make([]*Account, 0)
	if len(accountIDs) == 0 {
		return res, nil
	}
	err := DB.WithContext(ctx).Where("id in ?", accountIDs).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetAccountByUserID(ctx context.Context, userID int64) ([]*Account, error) {
	res := make([]*Account, 0)
	err := DB.WithContext(ctx).Where("user_id = ?", userID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ChangeBalance(ctx context.Context, accountID int64, delta float64) error {
	var account Account
	if err := DB.WithContext(ctx).Where("id = ?", accountID).First(&account).Error; err != nil {
		return err
	}
	account.Balance += delta
	log.Println("delta", delta)
	return DB.Save(&account).Error
}
