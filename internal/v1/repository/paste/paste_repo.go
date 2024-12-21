package paste

import (
	"context"
	"time"

	"github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/helper/constant"
	"github.com/kadekchresna/pastely/internal/v1/model"
	"gorm.io/gorm"
)

type pasteRepo struct {
	DB config.DB
}

func NewPasteRepo(DB config.DB) PasteRepo {
	return &pasteRepo{
		DB: DB,
	}
}

func (r *pasteRepo) GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error) {

	var res Paste

	if err := r.DB.SlaveDB.Raw(`
		SELECT  p.shortlink,
				p.expired_at,
				p.status,
				p.paste_url,
				p.created_at FROM paste p WHERE p.shortlink = ? AND p.status = 'active'
	`, params.Shortlink).Scan(&res).Error; err != nil {
		return nil, err
	}

	return &model.Paste{
		Shortlink:                 res.Shortlink,
		ExpirationLengthInMinutes: int(res.ExpiredAt.Sub(res.CreatedAt).Minutes()),
		PasteURL:                  res.PasteURL,
		Status:                    res.Status,
		CreatedAt:                 res.CreatedAt,
	}, nil
}

func (r *pasteRepo) CreatePaste(ctx context.Context, data model.Paste) (*model.Paste, error) {

	p := Paste{
		PasteURL: data.PasteURL,
	}

	trx := r.DB.MasterDB.Begin()
	p.Shortlink = ``
	p.Status = `active`

	if err := trx.Create(&p).Error; err != nil {
		trx.Rollback()
		return nil, err
	}

	p.GenerateShortURLBase62()

	returning := model.Paste{
		Shortlink:                 p.Shortlink,
		PasteURL:                  p.PasteURL,
		CreatedAt:                 p.CreatedAt,
		ExpiredAt:                 p.CreatedAt.Add(time.Duration(data.ExpirationLengthInMinutes) * time.Minute),
		ExpirationLengthInMinutes: data.ExpirationLengthInMinutes,
		Status:                    p.Status,
	}

	if err := trx.Model(p).Updates(map[string]interface{}{"shortlink": p.Shortlink, "expired_at": p.CreatedAt.Add(time.Duration(data.ExpirationLengthInMinutes) * time.Minute)}).Error; err != nil {
		trx.Rollback()
		return nil, err
	}

	if err := trx.Commit().Error; err != nil {
		trx.Rollback()
		return nil, err
	}

	return &returning, nil
}

func (r *pasteRepo) GetExpiredPastes(ctx context.Context) ([]model.Paste, error) {

	var res []Paste
	// err := r.DB.SlaveDB.Raw(`
	// 	SELECT  p.shortlink,
	// 			p.paste_url,
	// 			p.created_at
	// 			FROM paste p
	// 			WHERE p.expired_at < NOW()
	// `).Scan(&res).Error
	// if err != nil {
	// 	return nil, err
	// }

	if err := r.DB.SlaveDB.Where("expired_at < ?", time.Now()).Select("shortlink, paste_url, created_at").Find(&res).Error; err != nil {
		return nil, err
	}

	var returning []model.Paste
	for _, v := range res {
		returning = append(returning, model.Paste{
			PasteURL:  v.PasteURL,
			Shortlink: v.Shortlink,
		})
	}

	return returning, nil
}

func (r *pasteRepo) DeleteExpiredPastes(ctx context.Context, shortLinks []string) error {

	tx := r.DB.MasterDB.Begin()
	if _, ok := ctx.Value(constant.DB_TRANSACTION).(*gorm.DB); ok {
		tx = ctx.Value(constant.DB_TRANSACTION).(*gorm.DB)
	}

	if err := tx.Exec(`
		UPDATE paste SET status = 'deleted' WHERE shortlink IN (?)
	`, shortLinks).Error; err != nil {
		return err
	}

	return nil
}
