package paste

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
)

type PasteUsecase interface {
	GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error)
	CreatePaste(ctx context.Context, data CreatePaste) (*model.Paste, error)
	DeleteExpiredPastes(ctx context.Context) error
}
