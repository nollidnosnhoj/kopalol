package storage

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/nollidnosnhoj/kopalol/internal/config"
)

type S3Storage struct {
	bucket string
	client *s3.Client
	url    string
}

func NewS3Storage(context context.Context, config *config.Config) (*S3Storage, error) {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: config.S3_ENDPOINT,
		}, nil
	})

	cfg, err := awsConfig.LoadDefaultConfig(context,
		awsConfig.WithEndpointResolverWithOptions(r2Resolver),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.S3_ACCESS_KEY, config.S3_SECRET_KEY, "")),
		awsConfig.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		if config.S3_FORCE_PATH_STYLE {
			o.UsePathStyle = true
		}
	})
	return &S3Storage{
		bucket: config.UPLOAD_BUCKET_NAME,
		client: client,
		url:    config.S3_IMAGE_URL,
	}, nil
}

func (s *S3Storage) GetImageDir(filename string) string {
	return fmt.Sprintf("%s/%s", s.url, filename)
}

func (s *S3Storage) Get(filename string, context context.Context) (ImageResult, bool, error) {
	output, err := s.client.GetObject(context, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		if isNotFoundError(err) {
			return ImageResult{}, false, nil
		} else {
			return ImageResult{}, false, err
		}
	}
	byteArr, err := io.ReadAll(output.Body)
	if err != nil {
		return ImageResult{}, false, err
	}
	buffer := bytes.NewBuffer(byteArr)
	contentType := output.ContentType
	return ImageResult{Body: *buffer, ContentType: *contentType}, true, nil
}

func (s *S3Storage) Upload(context context.Context, filename string, contentType string, source io.Reader) error {
	_, err := s.client.PutObject(context, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(filename),
		Body:        source,
		ContentType: &contentType,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *S3Storage) Delete(context context.Context, filename string) error {
	_, err := s.client.DeleteObject(context, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		return err
	}
	return nil
}

func isNotFoundError(err error) bool {
	var apiError smithy.APIError
	if errors.As(err, &apiError) && apiError.ErrorCode() == "NoSuchKey" {
		return true
	}
	return false
}
