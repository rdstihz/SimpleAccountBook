package db

import (
	"context"
	"github.com/rdstihz/SimpleAccountBook/pkg/constants"
	"gorm.io/gorm"
)

type Bill struct {
	gorm.Model
	UserId    int64   `json:"user_id"`
	AccountId int64   `json:"account_id"`
	Type      int64   `json:"type"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	Comment   string  `json:"comment"`
}

func (b *Bill) TableName() string {
	return constants.BillTableName
}

func CreateBill(ctx context.Context, bills []*Bill) error {
	return DB.WithContext(ctx).Create(bills).Error
}

func DeleteBill(ctx context.Context, billID int64) error {
	return DB.WithContext(ctx).Delete(&Bill{}, billID).Error
}

func UpdateBill(ctx context.Context, billID int64, bill *Bill) error {
	return DB.WithContext(ctx).Model(&Bill{Model: gorm.Model{ID: uint(billID)}}).Updates(bill).Error
}

func MGetBill(ctx context.Context, billIDs []int64) ([]*Bill, error) {
	res := make([]*Bill, 0)
	if len(billIDs) == 0 {
		return res, nil
	}
	err := DB.WithContext(ctx).Where("id in ?", billIDs).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetBillByAccountID(ctx context.Context, accountID int64) ([]*Bill, error) {
	res := make([]*Bill, 0)
	err := DB.WithContext(ctx).Where("account_id = ?", accountID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
