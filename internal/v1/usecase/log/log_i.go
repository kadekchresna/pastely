package log

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
)

type LogUsecase interface {
	GetLog(ctx context.Context, params GetLogParams) (*model.Log, error)
	CreatePaste(ctx context.Context, data CreateLog) error
}
