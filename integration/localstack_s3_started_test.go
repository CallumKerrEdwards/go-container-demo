package integration

import (
	"net/http"
	"testing"
)

func TestLocalstackS3Started(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	const s3Endpoint = "http://localhost:4572"

	//when
	//get method is send to s3 localstack endpoint
	resp, err := http.Get(s3Endpoint)
	if err != nil {
		t.Error(err)
	}

	//then
	//S3 API is available at the given endpoint
	if resp.StatusCode != 200 {
		t.Errorf("S3 Endpoint %s was not available:\n Status code was %v",
			s3Endpoint, resp.StatusCode)
	}
}
