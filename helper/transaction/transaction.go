package transaction

import (
	"context"

	"github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/helper/constant"
	"gorm.io/gorm"
)

type transactionRepo struct {
	DB config.DB
}

func NewTransactionRepo(DB config.DB) TransactionRepo {
	return &transactionRepo{
		DB: DB,
	}
}

func (t *transactionRepo) TransactionWrapper(ctx context.Context, fn func(context.Context) error) error {

	tx := t.DB.MasterDB.Begin()
	if _, ok := ctx.Value(constant.DB_TRANSACTION).(*gorm.DB); ok {
		tx = ctx.Value(constant.DB_TRANSACTION).(*gorm.DB)
	}

	ctx = context.WithValue(ctx, constant.DB_TRANSACTION, tx)

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}

		tx.Commit()
	}()

	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
