package integration

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/callumkerredwards/localstack-go-wrapper/localstack"
	"github.com/callumkerredwards/localstack-go-wrapper/localstack/services"
	"github.com/stretchr/testify/assert"
)

var testCredentials = credentials.NewStaticCredentials("AKID", "SECRET", "SESSION")
var testRegion = aws.String("eu-west-1")
var disableSSL = aws.Bool(true)

func TestMain(m *testing.M) {
	os.Exit(testMainWrapper(m))
}

func testMainWrapper(m *testing.M) int {
	if !testing.Short() {
		s3Config := &services.ServiceConfig{
			Service: services.S3,
		}
		container, err := localstack.New(s3Config)
		if err != nil {
			log.Printf("Cannot create localstack, %v", err)
			return 1
		}
		err = container.Start()
		if err != nil {
			log.Printf("Cannot start localstack, %v", err)
			return 1
		}
		defer container.Stop()
	}
	return m.Run()
}

func TestLocalstackS3Started(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	//AWS session
	sess, err := session.NewSession(&aws.Config{
		Credentials: testCredentials,
		Region:      testRegion,
		Endpoint:    aws.String("http://localhost:4572"),
		DisableSSL:  disableSSL,
	})
	if err != nil {
		t.Errorf("Could not create new AWS session: %v", err)
	}
	s, err := json.MarshalIndent(sess, "", "  ")
	if err == nil {
		log.Printf("AWS Session is %s", string(s))
	}

	//when
	//Create S3 service client and call the list bucket API
	s3Client := s3.New(sess)
	_, err = s3Client.ListBuckets(nil)

	//then
	//API did not produce an error
	assert.Nil(t, err, "ListBuckets API should not produce an error")
}
