package integration

import (
	"bytes"
	"net/http"
	"testing"
	"time"
)

const S3Endpoint = "http://localhost:4572"

func TestLocalstackS3Started(t *testing.T) {
	if testing.Short() {
		t.Skip("Integration tests are skipped when running tests in short mode.")
	}

	//given
	// set up of http client
	client := http.Client{
		Timeout: time.Second * 2,
	}

	//when
	//get method is send to s3 localstack endpoint
	req, err := http.NewRequest(http.MethodGet, S3Endpoint, nil)
	if err != nil {
		t.Error(err)
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
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
