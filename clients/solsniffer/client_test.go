package solsniffer

import (
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetToken(t *testing.T) {
	type fields struct {
		apiKey string
		client httpclient.RetryableHttpClient
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "live test",
			fields: fields{apiKey: "bnhd3vwrqaodhmxwn4ry27zwxi9rnl"},
			args:   args{address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey)
			got, err := c.GetToken(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEmpty(t, got)
		})
	}
}
