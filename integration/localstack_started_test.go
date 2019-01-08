package integration

import (
	"os"
	"testing"

	"github.com/CallumKerrEdwards/go/container/receipt/localstack"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func TestMain(m *testing.M) {
	container := localstack.Start()

	result := m.Run()

	localstack.Stop(container)

	os.Exit(result)
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
	if err != nil {
		t.Errorf("Unable to list buckets, %v", err)
	}
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
	if err != nil {
		t.Errorf("Unable to list queues, %v", err)
	}
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
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
