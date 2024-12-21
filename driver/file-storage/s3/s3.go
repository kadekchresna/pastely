package s3

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"time"

	cfg "github.com/kadekchresna/pastely/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
)

type resolverV2 struct {
}

func (*resolverV2) ResolveEndpoint(ctx context.Context, params s3.EndpointParameters) (smithyendpoints.Endpoint, error) {
	if os.Getenv("APP_ENV") != cfg.PRODUCTION {
		u, err := url.Parse("http://localhost:9000")
		if err != nil {
			return smithyendpoints.Endpoint{}, err
		}
		return smithyendpoints.Endpoint{
			URI: *u,
		}, nil
	}

	// delegate back to the default v2 resolver otherwise
	return s3.NewDefaultEndpointResolverV2().ResolveEndpoint(ctx, params)
}

func CreateS3Client(ctx context.Context, c cfg.Config) (*S3BucketBasics, error) {

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(c.S3Region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	// If credentials are provided, use them; otherwise, use AWS default credential chain
	if c.S3AccessKey != "" && c.S3SecretKey != "" {
		cfg.Credentials = credentials.NewStaticCredentialsProvider(
			c.S3AccessKey,
			c.S3SecretKey,
			"", // session key
		)
	}

	return &S3BucketBasics{
		S3Client: s3.NewFromConfig(cfg),
	}, nil
}

func CreateMinIOClient(ctx context.Context, c cfg.Config) (*S3BucketBasics, error) {

	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(c.S3Region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(c.S3AccessKey, c.S3SecretKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return &S3BucketBasics{
		S3Client: s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
			o.BaseEndpoint = aws.String(c.S3Endpoint)
			// o.EndpointResolverV2 = &resolverV2{}
		}),
	}, nil
}

type S3BucketBasics struct {
	S3Client *s3.Client
}

func (basics *S3BucketBasics) UploadFile(ctx context.Context, bucketName string, objectKey string, content io.Reader, expires int) error {

	_, err := basics.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(objectKey),
		Body:        content,
		ContentType: aws.String("text/plain"),
		// Tagging:     aws.String("app=pastely"),
		Expires: aws.Time(time.Now().Add(time.Duration(expires+1) * time.Minute)), // add 1 min for threshold
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "EntityTooLarge" {
			log.Printf("Error while uploading object to %s. The object is too large.\n"+
				"To upload objects larger than 5GB, use the S3 console (160GB max)\n"+
				"or the multipart upload API (5TB max).", bucketName)
			return fmt.Errorf("error while uploading object to %s. The object is too large. To upload objects larger than 5GB, use the S3 console (160GB max) or the multipart upload API (5TB max)", bucketName)
		} else {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				content, bucketName, objectKey, err)
			return fmt.Errorf("couldn't upload file %v to %v:%v. Here's why: %v",
				content, bucketName, objectKey, err)
		}
	}

	// err = s3.NewObjectExistsWaiter(basics.S3Client).Wait(
	// 	ctx, &s3.HeadObjectInput{Bucket: aws.String(bucketName), Key: aws.String(objectKey)}, time.Minute)
	// if err != nil {
	// 	log.Printf("Failed attempt to wait for object %s to exist.\n", objectKey)
	// 	return fmt.Errorf("failed attempt to wait for object %s to exist", objectKey)

	// }

	return nil
}

func (basics *S3BucketBasics) ListBucket(ctx context.Context) error {
	count := 10
	fmt.Printf("Let's list up to %v buckets for your account.\n", count)
	result, err := basics.S3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
	if err != nil {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "AccessDenied" {
			fmt.Println("You don't have permission to list buckets for this account.")
		} else {
			fmt.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
		}
		return nil
	}
	if len(result.Buckets) == 0 {
		fmt.Println("You don't have any buckets!")
	} else {
		if count > len(result.Buckets) {
			count = len(result.Buckets)
		}
		for _, bucket := range result.Buckets[:count] {
			fmt.Printf("\t%v\n", *bucket.Name)
		}
	}

	return nil
}

func (basics *S3BucketBasics) GetFile(ctx context.Context, bucketName string, objectKey string) (string, error) {

	out, err := basics.S3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		var apiErr smithy.APIError
		if errors.As(err, &apiErr) && apiErr.ErrorCode() == "NoSuchKey" {
			log.Printf("Error while getting object from %s. The object does not exist.\n", bucketName)
			return "", fmt.Errorf("error while getting object from %s. The object does not exist", bucketName)
		} else {
			log.Printf("Couldn't get file %v from %v:%v. Here's why: %v\n",
				objectKey, bucketName, objectKey, err)
			return "", fmt.Errorf("couldn't get file %v from %v:%v. Here's why: %v",
				objectKey, bucketName, objectKey, err)
		}
	}

	defer out.Body.Close()

	content, err := io.ReadAll(out.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read file %v from %v:%v. Here's why: %v",
			objectKey, bucketName, objectKey, err)
	}

	return string(content), nil
}

func (basics *S3BucketBasics) DeleteFiles(ctx context.Context, bucketName string, objectKey []string) error {

	etag := ""
	identifiers := make([]types.ObjectIdentifier, 0, len(objectKey))
	for _, key := range objectKey {
		identifiers = append(identifiers, types.ObjectIdentifier{
			Key:  aws.String(key),
			ETag: aws.String(etag),
		})
	}

	_, err := basics.S3Client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &types.Delete{
			Objects: identifiers,
			Quiet:   aws.Bool(true),
		},
	})
	if err != nil {
		return fmt.Errorf("couldn't delete file %v from %v:%v. Here's why: %v",
			objectKey, bucketName, objectKey, err)
	}

	return nil
}
