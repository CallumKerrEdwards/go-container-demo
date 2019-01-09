package integration

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/callumkerredwards/receipt/localstack"
	"github.com/callumkerredwards/receipt/write"
	"github.com/stretchr/testify/assert"
)

func TestWriteToS3CreatesBucket(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	bucket := "test-bucket"
	sess := localstack.S3Session()
	s3Client := s3.New(sess)
	assert.NotContains(t, getListOfBuckets(s3Client), bucket)
	tmpFile := writeTempFile()

	//when
	err := write.ToS3(sess, bucket, tmpFile)

	//then
	if assert.NoError(t, err) {
		assert.Contains(t, getListOfBuckets(s3Client), bucket)
	}
}

func TestWriteToS3CreatesBucketFails(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	bucket := "test-bucket"
	tmpFile := writeTempFile()

	//when
	err := write.ToS3(localstack.SNSSession(), bucket, tmpFile)

	//then
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), bucket)
	}
}

func writeTempFile() *os.File {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Panicf("Cannot create temporary file %v", err)
	}

	// Remember to clean up the file afterwards
	defer os.Remove(tmpFile.Name())

	fmt.Println("Created File: " + tmpFile.Name())

	text := []byte("This is a golangcode.com example!")
	if _, err = tmpFile.Write(text); err != nil {
		log.Panicf("Failed to write to temporary file %v", err)
	}

	// Close the file
	if err := tmpFile.Close(); err != nil {
		log.Panic(err)
	}
	return tmpFile
}

func getListOfBuckets(s3Client *s3.S3) []string {
	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		log.Panicf("Unable to list buckets, %v", err)
	}

	names := make([]string, 0, len(result.Buckets))
	for _, b := range result.Buckets {
		names = append(names, aws.StringValue(b.Name))
	}
	return names
}
