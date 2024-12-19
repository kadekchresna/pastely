package paste

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
)

type PasteRepo interface {
	GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error)
}
