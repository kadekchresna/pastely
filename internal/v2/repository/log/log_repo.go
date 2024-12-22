package log

import (
	"context"
	"time"

	"github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/internal/v2/model"
)

type logRepo struct {
	DB config.DB
}

func NewLogRepo(DB config.DB) LogRepo {
	return &logRepo{
		DB: DB,
	}
}

func (l *logRepo) CreateLog(ctx context.Context, data model.Log) error {

	log := Log{
		Shortlink: data.Shortlink,
		Time:      time.Now(),
	}

	if err := l.DB.AnalyticDB.Create(&log).Error; err != nil {
		return err
	}

	return nil
}

func (l *logRepo) GetLog(ctx context.Context, params GetLogParams) (*model.Log, error) {

	var res LogAggregate

	db := l.DB.AnalyticDB.Select("shortlink, COUNT(shortlink) as total").Group("shortlink").Table("paste_log")

	if !(params.DateStart.IsZero() && params.DateEnd.IsZero()) {
		db = db.Where("time >= ? AND time <= ?", params.DateStart, params.DateEnd)
	}

	db.Where("shortlink = ?", params.Shortlink)

	if err := db.Find(&res).Error; err != nil {
		return nil, err
	}

	return &model.Log{
		Shortlink: params.Shortlink,
		Total:     res.Total,
		DateStart: params.DateStart,
		DateEnd:   params.DateEnd,
	}, nil

}
