package transaction

import "context"

type TransactionRepo interface {
	TransactionWrapper(ctx context.Context, fn func(context.Context) error) error
}
