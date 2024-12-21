package paste

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/kadekchresna/pastely/config"
	filestorage "github.com/kadekchresna/pastely/driver/file-storage"
	"github.com/kadekchresna/pastely/helper/transaction"
	"github.com/kadekchresna/pastely/internal/v1/model"
	"github.com/kadekchresna/pastely/internal/v1/repository/paste"
)

type pasteUsecase struct {
	Config      config.Config
	PasteRepo   paste.PasteRepo
	Transaction transaction.TransactionRepo
	FileStorage filestorage.Bucket
}

func NewPasteUsecase(Config config.Config, PasteRepo paste.PasteRepo, Transaction transaction.TransactionRepo) PasteUsecase {
	return &pasteUsecase{
		Config:      Config,
		PasteRepo:   PasteRepo,
		Transaction: Transaction,
		FileStorage: filestorage.NewBucket(Config.AppFileStorage, Config),
	}
}

func (u *pasteUsecase) GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error) {

	paste, err := u.PasteRepo.GetPaste(ctx, paste.NewGetPasteParams(params.Shortlink))
	if err != nil {
		return nil, err
	}

	content, err := u.FileStorage.GetFile(ctx, u.Config.S3BucketName, paste.PasteURL)
	if err != nil {
		return nil, err
	}

	paste.PasteContent = content
	return paste, nil
}

func (u *pasteUsecase) CreatePaste(ctx context.Context, data CreatePaste) (*model.Paste, error) {

	now := time.Now()
	reader := bytes.NewReader([]byte(data.PasteContent))
	objectKey := fmt.Sprintf("%s-%d.txt", now.Format("2006-01-02"), now.Unix())
	if err := u.FileStorage.UploadFile(ctx, u.Config.S3BucketName, objectKey, reader, data.ExpirationLengthInMinutes); err != nil {
		return nil, err
	}

	// write the link to db master
	res, err := u.PasteRepo.CreatePaste(ctx, model.Paste{ExpirationLengthInMinutes: data.ExpirationLengthInMinutes, PasteURL: objectKey})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *pasteUsecase) DeleteExpiredPastes(ctx context.Context) error {

	expiredPastes, err := u.PasteRepo.GetExpiredPastes(ctx)
	if err != nil {
		return err
	}

	if len(expiredPastes) == 0 {
		return nil
	}

	shortLinks := make([]string, 0)
	objectKeys := make([]string, 0)
	for _, paste := range expiredPastes {
		shortLinks = append(shortLinks, paste.Shortlink)
		objectKeys = append(objectKeys, paste.PasteURL)
	}

	if err := u.Transaction.TransactionWrapper(ctx, func(ctxInside context.Context) error {

		if err := u.PasteRepo.DeleteExpiredPastes(ctxInside, shortLinks); err != nil {
			return err
		}

		if err := u.FileStorage.DeleteFiles(ctxInside, u.Config.S3BucketName, objectKeys); err != nil {
			return err
		}

		return nil

	}); err != nil {
		return err
	}

	return nil
}
