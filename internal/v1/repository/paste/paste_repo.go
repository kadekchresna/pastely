package paste

import (
	"context"

	"github.com/kadekchresna/pastely/internal/v1/model"
	"gorm.io/gorm"
)

type pasteRepo struct {
	DB *gorm.DB
}

func NewPasteRepo(DB *gorm.DB) PasteRepo {
	return &pasteRepo{
		DB: DB,
	}
}

func (r *pasteRepo) GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error) {

	var res Paste

	if err := r.DB.Raw(`
		select shortlink
			expiration_length_in_minutes
			paste_url
			created_at from paste p
	`).Where(`shortlink = ?`, params.Shortlink).Scan(&res).Error; err != nil {
		return nil, err
	}

	p := model.Paste(res)

	return &p, nil
}
