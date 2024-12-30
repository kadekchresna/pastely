package paste

import (
	"context"
	"fmt"

	"github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/driver/cache"
	filestorage "github.com/kadekchresna/pastely/driver/file-storage"
	"github.com/kadekchresna/pastely/helper/constant"
	"github.com/kadekchresna/pastely/helper/transaction"
	"github.com/kadekchresna/pastely/internal/v2/model"
	"github.com/kadekchresna/pastely/internal/v2/repository/log"
	"github.com/kadekchresna/pastely/internal/v2/repository/paste"
)

type pasteUsecase struct {
	Config      config.Config
	PasteRepo   paste.PasteRepo
	Transaction transaction.TransactionRepo
	LogRepo     log.LogRepo
	Cache       cache.Cache

	FileStorage filestorage.Bucket
}

func NewPasteUsecase(Config config.Config, PasteRepo paste.PasteRepo, Transaction transaction.TransactionRepo, LogRepo log.LogRepo, Cache cache.Cache) PasteUsecase {
	return &pasteUsecase{
		Config:      Config,
		PasteRepo:   PasteRepo,
		Transaction: Transaction,
		LogRepo:     LogRepo,
		FileStorage: filestorage.NewBucket(Config.AppFileStorage, Config),
		Cache:       Cache,
	}
}

func (u *pasteUsecase) GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error) {

	var p *model.Paste
	var err error

	key := fmt.Sprintf("%s-%s", constant.KEY_CACHE_DETAIL_PASTE_USECASE, params.Shortlink)
	err = u.Cache.Get(ctx, key, p)
	if p != nil {
		return p, nil
	}

	p, err = u.PasteRepo.GetPaste(ctx, paste.NewGetPasteParams(params.Shortlink))
	if err != nil {
		return nil, err
	}

	if err := u.LogRepo.CreateLog(ctx, model.Log{Shortlink: p.Shortlink}); err != nil {
		return nil, err
	}

	go u.Cache.Set(context.Background(), key, 60, p)

	return p, nil
}

func (u *pasteUsecase) CreatePaste(ctx context.Context, data CreatePaste) (*model.Paste, error) {

	// write the link to db master
	res, err := u.PasteRepo.CreatePaste(ctx, model.Paste{ExpirationLengthInMinutes: data.ExpirationLengthInMinutes, PasteURL: data.PasteURL})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *pasteUsecase) GetPresignedURL(ctx context.Context, objectKey string, expires int) (*filestorage.PresignedHTTPResponse, error) {

	var f *filestorage.PresignedHTTPResponse
	var err error

	key := fmt.Sprintf("%s-%s", constant.KEY_CACHE_PRESIGNED_URL_PASTE_USECASE, objectKey)
	err = u.Cache.Get(ctx, key, f)
	if f != nil {
		return f, nil
	}

	f, err = u.FileStorage.GenerateGetPresignedURL(ctx, u.Config.S3BucketName, objectKey, expires)
	if err != nil {
		return nil, err
	}

	go u.Cache.Set(context.Background(), key, expires*50, f)

	return f, nil
}

func (u *pasteUsecase) PutPresignedURL(ctx context.Context, objectKey string, expires int) (*filestorage.PresignedHTTPResponse, error) {
	req, err := u.FileStorage.GeneratePutPresignedURL(ctx, u.Config.S3BucketName, objectKey, expires)
	if err != nil {
		return nil, err
	}

	return req, nil
}
