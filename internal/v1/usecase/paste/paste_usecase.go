package paste

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
	"github.com/kadekchresna/pastely/internal/v1/repository/paste"
)

type pasteUsecase struct {
	PasteRepo paste.PasteRepo
}

func NewPasteUsecase(PasteRepo paste.PasteRepo) PasteUsecase {
	return &pasteUsecase{
		PasteRepo: PasteRepo,
	}
}

func (u *pasteUsecase) GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error) {

	paste, err := u.PasteRepo.GetPaste(ctx, paste.NewGetPasteParams(params.Shortlink))
	if err != nil {
		return nil, err
	}

	return paste, nil
}
