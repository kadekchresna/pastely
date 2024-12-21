package paste

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
)

type PasteRepo interface {
	GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error)
	CreatePaste(ctx context.Context, data model.Paste) (*model.Paste, error)
	GetExpiredPastes(ctx context.Context) ([]model.Paste, error)
	DeleteExpiredPastes(ctx context.Context, shortLinks []string) error
}
