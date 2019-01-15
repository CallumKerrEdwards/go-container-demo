package integration

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/callumkerredwards/receipt/write"
	"github.com/stretchr/testify/assert"
)

func TestWriteToS3CreatesBucket(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	bucket := "test-bucket"
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:4572"),
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}
	s3Client := s3.New(sess, &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
	})
	assert.NotContains(t, getListOfBuckets(t, s3Client), bucket)
	tmpFile := writeTempFile(t)
	defer os.Remove(tmpFile.Name())

	//when
	err = write.ToS3(sess, bucket, tmpFile.Name())
	defer cleanupBucket(t, s3Client, bucket)

	//then
	if assert.NoError(t, err) {
		assert.Contains(t, getListOfBuckets(t, s3Client), bucket)
	}
}

func TestWriteToS3UploadsFile(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	bucket := "test-bucket"
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:4572"),
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}
	s3Client := s3.New(sess, &aws.Config{
		S3ForcePathStyle: aws.Bool(true),
	})
	assert.NotContains(t, getListOfBuckets(t, s3Client), bucket)
	tmpFile := writeTempFile(t)
	defer os.Remove(tmpFile.Name())

	//when
	err = write.ToS3(sess, bucket, tmpFile.Name())
	defer cleanupBucket(t, s3Client, bucket)

	//then
	_, filename := filepath.Split(tmpFile.Name())
	if assert.NoError(t, err) {
		assert.Contains(t, getListOfObjectsInBuckets(t, s3Client, bucket), filename)
	}
}

func writeTempFile(t *testing.T) *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "testfile-")
	if err != nil {
		log.Panicf("Cannot create temporary file %v", err)
	}

	fmt.Println("Created File: " + tmpFile.Name())

	text := []byte("Transaction data")
	if _, err = tmpFile.Write(text); err != nil {
		log.Panicf("Failed to write to temporary file %v", err)
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		t.Error(err)
	}
	return tmpFile
}

func getListOfBuckets(t *testing.T, s3Client *s3.S3) []string {
	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		t.Errorf("Unable to list buckets, %v", err)
	}

	// turn slice of result buckets into a slice of bucket names
	names := make([]string, 0, len(result.Buckets))
	for _, b := range result.Buckets {
		names = append(names, aws.StringValue(b.Name))
	}
	return names
}

func getListOfObjectsInBuckets(t *testing.T, s3Client *s3.S3, bucket string) []string {
	result, err := s3Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucket)})
	if err != nil {
		t.Errorf("Unable to list objects in bucket %q, %v", bucket, err)
	}

	// turn slice of result objects into a slice of object names
	names := make([]string, 0, len(result.Contents))
	for _, b := range result.Contents {
		names = append(names, aws.StringValue(b.Key))
	}
	return names
}

func cleanupBucket(t *testing.T, s3Client *s3.S3, bucket string) {
	for _, s := range getListOfObjectsInBuckets(t, s3Client, bucket) {
		log.Printf("Deleting %v from bucket %v", s, bucket)
		_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(s)})
		if err != nil {
			t.Errorf("Unable to delete object %q from bucket %q, %v", s, bucket, err)
		}
	}

	_, err := s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		t.Errorf("Unable to delete  bucket %q, %v", bucket, err)
	}

	err = s3Client.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		t.Errorf("Error occurred while waiting for bucket to be deleted, %v", bucket)
	}

}
