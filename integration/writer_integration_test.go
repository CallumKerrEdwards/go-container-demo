package integration

import (
	"os"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/callumkerredwards/receipt/localstack"
	"github.com/callumkerredwards/receipt/write"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	var container string
	if !testing.Short() {
		container = localstack.Start()
	}

	result := m.Run()

	if !testing.Short() {
		localstack.Stop(container)
	}

	os.Exit(result)
}

func TestWriteToS3CreatesBucket(t *testing.T) {

	//given
	bucket := "test-bucket"
	sess := localstack.S3Session()
	s3Client := s3.New(sess)
	assert.NotContains(t, getListOfBuckets(s3Client), bucket)

	//when
	write.ToS3(sess, bucket, "data")

	//then
	assert.Contains(t, getListOfBuckets(s3Client), bucket)

}

func getListOfBuckets(s3Client *s3.S3) []string {
    result, err := s3Client.ListBuckets(nil)
	if err != nil {
        exitErrorf("Unable to list buckets, %v", err)
    }

	names := make([]string, 0, len(result.Buckets))
	for _, b := range result.Buckets {
		names = append(names, aws.StringValue(b.Name))
	}
	return names
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}
