package minio

import (
	"context"
	"io"

	"github.com/kadekchresna/pastely/config"
	filestorage "github.com/kadekchresna/pastely/driver/file-storage"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioBucketBasics struct {
	filestorage.BucketBasics
	MinioClient *minio.Client
}

func CreateMinioClient(ctx context.Context, c config.Config) (*MinioBucketBasics, error) {

	minioClient, err := minio.New(c.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.S3AccessKey, c.S3SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &MinioBucketBasics{
		MinioClient: minioClient,
	}, nil
}

func (basics *MinioBucketBasics) UploadFile(ctx context.Context, bucketName string, objectKey string, content io.Reader) error {

	_, err := basics.MinioClient.PutObject(ctx, bucketName, objectKey, content, -1, minio.PutObjectOptions{
		ContentType: "text/plain",
		Internal:    minio.AdvancedPutOptions{},
	})
	if err != nil {
		return err
	}

	return nil
}
