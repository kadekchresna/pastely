package paste

import (
	"context"

	"github.com/kadekchresna/pastely/config"
	filestorage "github.com/kadekchresna/pastely/driver/file-storage"
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

	FileStorage filestorage.Bucket
}

func NewPasteUsecase(Config config.Config, PasteRepo paste.PasteRepo, Transaction transaction.TransactionRepo, LogRepo log.LogRepo) PasteUsecase {
	return &pasteUsecase{
		Config:      Config,
		PasteRepo:   PasteRepo,
		Transaction: Transaction,
		LogRepo:     LogRepo,
		FileStorage: filestorage.NewBucket(Config.AppFileStorage, Config),
	}
}

func (u *pasteUsecase) GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error) {

	paste, err := u.PasteRepo.GetPaste(ctx, paste.NewGetPasteParams(params.Shortlink))
	if err != nil {
		return nil, err
	}

	if err := u.LogRepo.CreateLog(ctx, model.Log{Shortlink: paste.Shortlink}); err != nil {
		return nil, err
	}

	return paste, nil
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
	req, err := u.FileStorage.GenerateGetPresignedURL(ctx, u.Config.S3BucketName, objectKey, expires)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (u *pasteUsecase) PutPresignedURL(ctx context.Context, objectKey string, expires int) (*filestorage.PresignedHTTPResponse, error) {
	req, err := u.FileStorage.GeneratePutPresignedURL(ctx, u.Config.S3BucketName, objectKey, expires)
	if err != nil {
		return nil, err
	}

	return req, nil
}
