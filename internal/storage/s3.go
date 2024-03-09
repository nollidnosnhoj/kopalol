package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/nollidnosnhoj/vgpx/internal/config"
)

type S3Storage struct {
	bucket string
	Client *s3.Client
}

func NewS3Storage(context context.Context, config *config.Config) (*S3Storage, error) {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", config.CLOUDFLARE_ACCOUNT_ID),
		}, nil
	})

	cfg, err := awsConfig.LoadDefaultConfig(context,
		awsConfig.WithEndpointResolverWithOptions(r2Resolver),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.CLOUDFLARE_ACCESS_KEY_ID, config.CLOUDFLARE_ACCESS_KEY_SECRET, "")),
		awsConfig.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)
	return &S3Storage{
		bucket: config.UPLOAD_BUCKET_NAME,
		Client: client,
	}, nil
}

func (s *S3Storage) Get(filename string, context context.Context) (ImageResult, error) {
	output, err := s.Client.GetObject(context, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return ImageResult{}, err
	}
	byteArr, err := io.ReadAll(output.Body)
	if err != nil {
		return ImageResult{}, err
	}
	buffer := bytes.NewBuffer(byteArr)
	contentType := output.ContentType
	return ImageResult{Body: *buffer, ContentType: *contentType}, nil
}

func (s *S3Storage) Upload(filename string, source io.Reader, context context.Context) error {
	_, err := s.Client.PutObject(context, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
		Body:   source,
	})
	if err != nil {
		return err
	}
	return nil
}