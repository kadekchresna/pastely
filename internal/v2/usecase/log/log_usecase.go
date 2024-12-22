package log

import (
	"context"

	"github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/internal/v2/model"
	"github.com/kadekchresna/pastely/internal/v2/repository/log"
)

type logUsecase struct {
	Config  config.Config
	LogRepo log.LogRepo
}

func NewLogUsecase(Config config.Config, LogRepo log.LogRepo) LogUsecase {
	return &logUsecase{
		Config:  Config,
		LogRepo: LogRepo,
	}
}

func (u *logUsecase) GetLog(ctx context.Context, params GetLogParams) (*model.Log, error) {

	log, err := u.LogRepo.GetLog(ctx, log.NewGetLogParams(params.Shortlink, params.DateStart, params.DateEnd))
	if err != nil {
		return nil, err
	}

	return log, nil
}

func (u *logUsecase) CreatePaste(ctx context.Context, data CreateLog) error {

	return u.LogRepo.CreateLog(ctx, model.Log{Shortlink: data.Shortlink})
}
