package paste

import (
	"context"

	filestorage "github.com/kadekchresna/pastely/driver/file-storage"
	"github.com/kadekchresna/pastely/internal/v2/model"
)

type PasteUsecase interface {
	GetPaste(ctx context.Context, params GetPasteParams) (*model.Paste, error)
	CreatePaste(ctx context.Context, data CreatePaste) (*model.Paste, error)
	GetPresignedURL(ctx context.Context, objectKey string, expires int) (*filestorage.PresignedHTTPResponse, error)
	PutPresignedURL(ctx context.Context, objectKey string, expires int) (*filestorage.PresignedHTTPResponse, error)
}
