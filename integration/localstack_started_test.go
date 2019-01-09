package integration

import (
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/callumkerredwards/receipt/localstack"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	if !testing.Short() {
		container, err := localstack.Start()
		if err != nil {
			log.Panicf("Cannot start localstack, %v", err)
		}
		defer localstack.Stop(container)
	}
	return m.Run()
}

func TestLocalstackS3Started(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess := localstack.S3Session()

	//when
	//Create S3 service client and call the list bucket API
	s3Client := s3.New(sess)
	_, err := s3Client.ListBuckets(nil)

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListBuckets API should not produce an error")
}

func TestLocalstackSQSStarted(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess := localstack.SQSSession()

	//when
	//Create SQS client cand call the list queues API
	sqsClient := sqs.New(sess)
	_, err := sqsClient.ListQueues(nil)

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListQueues API should not produce an error")
}

func TestLocalstackSNSStarted(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess := localstack.SNSSession()

	//when
	//Create SNS client and call the list topics API
	snsClient := sns.New(sess)
	_, err := snsClient.ListTopics(&sns.ListTopicsInput{})

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListTopics API should not produce an error")
}
