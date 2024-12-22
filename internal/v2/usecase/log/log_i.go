package log

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v2/model"
)

type LogUsecase interface {
	GetLog(ctx context.Context, params GetLogParams) (*model.Log, error)
	CreatePaste(ctx context.Context, data CreateLog) error
}
