package filestorage

import (
	"context"
	"io"
	"net/http"

	"github.com/kadekchresna/pastely/config"
	"github.com/kadekchresna/pastely/driver/file-storage/s3"
)

const (
	S3    = `s3`
	MINIO = `minio`
)

type BucketBasics struct {
	BucketTarget string
	cfg          config.Config
}

type PresignedHTTPResponse struct {
	URL          string      `json:"url"`
	Method       string      `json:"method"`
	SignedHeader http.Header `json:"signed_header"`
}
type Bucket interface {
	UploadFile(ctx context.Context, bucketName string, objectKey string, content io.Reader, expires int) error
	GetFile(ctx context.Context, bucketName string, objectKey string) (string, error)
	DeleteFiles(ctx context.Context, bucketName string, objectKey []string) error
	GenerateGetPresignedURL(ctx context.Context, bucketName string, objectKey string, expires int) (*PresignedHTTPResponse, error)
	GeneratePutPresignedURL(ctx context.Context, bucketName string, objectKey string, expires int) (*PresignedHTTPResponse, error)
}

func NewBucket(bucketTarget string, cfg config.Config) *BucketBasics {
	return &BucketBasics{
		BucketTarget: bucketTarget,
		cfg:          cfg,
	}
}

func (b *BucketBasics) UploadFile(ctx context.Context, bucketName string, objectKey string, content io.Reader, expires int) error {

	switch b.BucketTarget {
	case MINIO:
		client, err := s3.CreateMinIOClient(ctx, b.cfg)
		if err != nil {
			return err
		}

		if err := client.UploadFile(ctx, b.cfg.S3BucketName, objectKey, content, expires); err != nil {
			return err
		}

	default:
		client, err := s3.CreateS3Client(ctx, b.cfg)
		if err != nil {
			return err
		}

		if err := client.UploadFile(ctx, b.cfg.S3BucketName, objectKey, content, expires); err != nil {
			return err
		}

	}

	return nil
}

func (b *BucketBasics) GetFile(ctx context.Context, bucketName string, objectKey string) (string, error) {

	switch b.BucketTarget {
	case MINIO:
		client, err := s3.CreateMinIOClient(ctx, b.cfg)
		if err != nil {
			return "", err
		}

		return client.GetFile(ctx, b.cfg.S3BucketName, objectKey)

	default:
		client, err := s3.CreateS3Client(ctx, b.cfg)
		if err != nil {
			return "", err
		}

		return client.GetFile(ctx, b.cfg.S3BucketName, objectKey)
	}

}

func (b *BucketBasics) DeleteFiles(ctx context.Context, bucketName string, objectKey []string) error {

	switch b.BucketTarget {
	case MINIO:
		client, err := s3.CreateMinIOClient(ctx, b.cfg)
		if err != nil {
			return err
		}

		return client.DeleteFiles(ctx, b.cfg.S3BucketName, objectKey)
	default:
		client, err := s3.CreateS3Client(ctx, b.cfg)
		if err != nil {
			return err
		}

		return client.DeleteFiles(ctx, b.cfg.S3BucketName, objectKey)
	}
}

func (b *BucketBasics) GenerateGetPresignedURL(ctx context.Context, bucketName string, objectKey string, expires int) (*PresignedHTTPResponse, error) {

	switch b.BucketTarget {
	case MINIO:
		client, err := s3.CreateMinIOClient(ctx, b.cfg)
		if err != nil {
			return nil, err
		}

		res, err := client.GetObjectPresignedURL(ctx, b.cfg.S3BucketName, objectKey, expires)

		return &PresignedHTTPResponse{
			URL:          res.URL,
			Method:       res.Method,
			SignedHeader: res.SignedHeader,
		}, err
	default:
		client, err := s3.CreateS3Client(ctx, b.cfg)
		if err != nil {
			return nil, err
		}
		res, err := client.GetObjectPresignedURL(ctx, b.cfg.S3BucketName, objectKey, expires)

		return &PresignedHTTPResponse{
			URL:          res.URL,
			Method:       res.Method,
			SignedHeader: res.SignedHeader,
		}, err
	}
}

func (b *BucketBasics) GeneratePutPresignedURL(ctx context.Context, bucketName string, objectKey string, expires int) (*PresignedHTTPResponse, error) {

	switch b.BucketTarget {
	case MINIO:
		client, err := s3.CreateMinIOClient(ctx, b.cfg)
		if err != nil {
			return nil, err
		}

		res, err := client.CreateObjectPresignedURL(ctx, b.cfg.S3BucketName, objectKey, expires)

		return &PresignedHTTPResponse{
			URL:          res.URL,
			Method:       res.Method,
			SignedHeader: res.SignedHeader,
		}, err
	default:
		client, err := s3.CreateS3Client(ctx, b.cfg)
		if err != nil {
			return nil, err
		}
		res, err := client.CreateObjectPresignedURL(ctx, b.cfg.S3BucketName, objectKey, expires)

		return &PresignedHTTPResponse{
			URL:          res.URL,
			Method:       res.Method,
			SignedHeader: res.SignedHeader,
		}, err
	}
}
