package integration

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var testCredentials = credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")

func TestLocalstackS3Started(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      aws.String("eu-west-1"),
		Endpoint:    aws.String("http://localhost:4572"),
		DisableSSL:  aws.Bool(true),
	})
	if err != nil {
		t.Errorf("Unable to create AWS session, %v", err)
	}

	//when
	//Create S3 service client
	s3Client := s3.New(sess)

	//then
	//S3 API is able to list buckets
	_, err2 := s3Client.ListBuckets(nil)
	if err2 != nil {
		t.Errorf("Unable to list buckets, %v", err2)
	}
}

func TestLocalstackSQSStarted(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      aws.String("eu-west-1"),
		Endpoint:    aws.String("http://localhost:4576"),
		DisableSSL:  aws.Bool(true),
	})
	if err != nil {
		t.Errorf("Unable to create AWS session, %v", err)
	}

	//when
	//Create SQS client
	sqsClient := sqs.New(sess)

	//then
	//SQS API is able to list queues
	_, err2 := sqsClient.ListQueues(nil)
	if err2 != nil {
		t.Errorf("Unable to list queues, %v", err2)
	}
}

func TestLocalstackSNSStarted(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      aws.String("eu-west-1"),
		Endpoint:    aws.String("http://localhost:4576"),
		DisableSSL:  aws.Bool(true),
	})
	if err != nil {
		t.Errorf("Unable to create AWS session, %v", err)
	}

	//when
	//Create SQS client
	snsClient := sns.New(sess)

	//then
	//SNS API is able to list topics
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	params := &sns.ListTopicsInput{}
	_, err2 := snsClient.ListTopicsWithContext(ctx, params)
	if err2 != nil {
		t.Errorf("expect no error, got %v", err2)
	}
}
