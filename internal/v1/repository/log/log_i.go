package log

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
)

type LogRepo interface {
	GetLog(ctx context.Context, params GetLogParams) (*model.Log, error)
	CreateLog(ctx context.Context, data model.Log) error
}
