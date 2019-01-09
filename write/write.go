package write

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
)

// ToS3 writes the provided file to S3 in the provided bucket
func ToS3(sess *session.Session, bucket string, file *os.File) error {

	fmt.Printf("Session %v, bucketName %v, file content %s\n", *sess.Config.Endpoint, bucket, file.Name())

	svc := s3.New(sess, &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
	})

	if err := createBucket(svc, bucket); err != nil {
		return errors.Wrapf(err, "Unable to create bucket %q", bucket)
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
