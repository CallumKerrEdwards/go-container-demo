package write

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

// ToS3 writes the provided file to S3 in the provided bucket
func ToS3(sess *session.Session, bucket string, filename string) error {

	svc := s3.New(sess, &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
	})

	if err := createBucket(svc, bucket); err != nil {
		return errors.Wrapf(err, "Unable to create bucket %q", bucket)
	}

	uploader := s3manager.NewUploaderWithClient(svc)
	if err := uploadFile(uploader, bucket, filename); err != nil {
		return errors.Wrapf(err, "Unable to upload file %q", filename)
	}
	return nil
}

// Create the S3 Bucket
func createBucket(svc *s3.S3, bucket string) error {
	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Bucket %q successfully created\n", bucket)
	return nil
}

// Upload the file
func uploadFile(uploader *s3manager.Uploader, bucket string, fpath string) error {
	file, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, name := filepath.Split(fpath)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
		Body:   file,
	})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded %q to %q\n", name, bucket)
	return nil
}
