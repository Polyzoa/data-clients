package honeypot

import (
	"encoding/json"
	"testing"
)

func Test_DataResponse(t *testing.T) {
	var response DataResponse
	err := json.Unmarshal([]byte(WeirdoResponseData), &response)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
