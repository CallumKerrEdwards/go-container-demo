package integration

import (
	"bytes"
	"net/http"
	"testing"
)

const S3Endpoint = "http://localhost:4572"

func TestLocalstackS3Started(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given

	//when
	//get method is send to s3 localstack endpoint
	resp, err := http.Get(S3Endpoint)
	if err != nil {
		t.Error(err)
	}

	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()

	//then
	//S3 API is available at the given endpoint
	if resp.StatusCode != 200 {
		t.Errorf("S3 Endpoint %s was not available:\n"+
			"Status code was %v, body was %s", S3Endpoint, resp.StatusCode, body)
	}
}
