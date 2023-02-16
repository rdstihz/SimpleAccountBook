package service

import (
	"context"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/dal/db"
	"github.com/rdstihz/SimpleAccountBook/cmd/account/pack"
	"github.com/rdstihz/SimpleAccountBook/kitex_gen/account"
	"github.com/rdstihz/SimpleAccountBook/pkg/errno"
)

// CreateBill 创建帐单
func CreateBill(ctx context.Context, bill *account.Bill) error {
	// 修改帐户余额
	if err := ChangeBalance(ctx, bill.AccountId, float64(bill.Type)*bill.Amount); err != nil {
		return err
	}
	// 保存帐单
	return db.CreateBill(ctx, []*db.Bill{
		&db.Bill{
			UserId:    bill.UserId,
			AccountId: bill.AccountId,
			Type:      bill.Type,
			Amount:    bill.Amount,
			Category:  bill.Category,
			Comment:   bill.Comment,
		},
	})
}

// DeleteBill 删除帐单
func DeleteBill(ctx context.Context, billID int64) error {
	// 修改帐户余额
	bill, err := MGetBill(ctx, []int64{billID})
	if err != nil {
		return err
	}
	if len(bill) == 0 {
		return errno.BillNotFoundErr
	}
	delta := -1.0 * float64(bill[0].Type) * bill[0].Amount
	err = ChangeBalance(ctx, bill[0].AccountId, delta)
	if err != nil {
		return err
	}

	// 从数据库中删除帐单
	return db.DeleteBill(ctx, billID)
}

// UpdateBill 更新帐单
func UpdateBill(ctx context.Context, bill *account.Bill) error {
	// 修改帐户余额
	oldBill, err := MGetBill(ctx, []int64{bill.BillId})
	if err != nil {
		return err
	}
	if len(oldBill) == 0 {
		return errno.BillNotFoundErr
	}
	oldAmount := float64(oldBill[0].Type) * oldBill[0].Amount
	newAmount := float64(bill.Type) * bill.Amount
	err = ChangeBalance(ctx, oldBill[0].AccountId, -1.0*oldAmount)
	if err != nil {
		return err
	}
	err = ChangeBalance(ctx, bill.AccountId, newAmount)
	if err != nil {
		return err
	}
	//保存到数据库
	return db.UpdateBill(ctx, bill.BillId, &db.Bill{
		UserId:    bill.UserId,
		AccountId: bill.AccountId,
		Type:      bill.Type,
		Amount:    bill.Amount,
		Category:  bill.Category,
		Comment:   bill.Comment,
	})
}

// MGetBill 查询帐单
func MGetBill(ctx context.Context, billIDs []int64) ([]*account.Bill, error) {
	model, err := db.MGetBill(ctx, billIDs)
	if err != nil {
		return nil, err
	}
	return pack.Bills(model), nil
}

// GetBillByAccountID 查询帐户下的所有帐单
func GetBillByAccountID(ctx context.Context, accountID int64) ([]*account.Bill, error) {
	model, err := db.GetBillByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}
	return pack.Bills(model), nil
}
