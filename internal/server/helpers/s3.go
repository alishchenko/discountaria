package helpers

import (
	"fmt"
	"github.com/alishchenko/discountaria/internal/config"
	"github.com/pkg/errors"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)
import "github.com/aws/aws-sdk-go/aws/session"

func NewAWSSession(config *config.AWSConfig) *session.Session {
	awsSessionConfig := aws.Config{
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyID,
			config.SecretKeyID,
			""),
		Region:           aws.String(config.Region),
		DisableSSL:       aws.Bool(config.SslDisable),
		S3ForcePathStyle: aws.Bool(config.ForcePathStyle),
	}

	if config.Endpoint != "" {
		awsSessionConfig.Endpoint = aws.String(config.Endpoint)
	}

	return session.Must(session.NewSession(&awsSessionConfig))
}

func UploadFile(file multipart.File, key string, config *config.AWSConfig) error {
	awsSession := NewAWSSession(config)
	uploader := s3manager.NewUploader(awsSession)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}

func GetUrl(key string, config *config.AWSConfig) (string, error) {
	// as the bucket is open, we can simply `glue` the link
	return fmt.Sprintf(
		"https://%s.s3.%s.amazonaws.com/%s",
		config.Bucket,
		config.Region,
		key,
	), nil
}

func DeleteFile(key string, config *config.AWSConfig) error {
	awsSession := NewAWSSession(config)
	service := s3.New(awsSession)

	_, err := service.DeleteObject(&s3.DeleteObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(config.Bucket),
	})
	if err != nil {
		return err
	}

	err = service.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(config.Bucket),
	})

	return err
}

func IsKeyExists(key string, config *config.AWSConfig) (bool, error) {
	awsSession := NewAWSSession(config)
	service := s3.New(awsSession)

	_, err := service.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound":
				return false, nil
			default:
				return false, err
			}
		}

		return false, err
	}

	return true, nil
}

func CheckDocumentMimeType(mtype string, r *http.Request) (string, error) {
	for _, el := range MimeTypes(r).AllowedMimeTypes {
		if el == mtype {
			return strings.Split(mtype, "/")[1], nil
		}
	}
	return "", errors.New("invalid file extension")
}
