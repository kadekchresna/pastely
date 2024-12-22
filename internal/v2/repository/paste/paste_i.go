package paste

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v2/model"
)

type PasteRepo interface {
	GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error)
	CreatePaste(ctx context.Context, data model.Paste) (*model.Paste, error)
}
