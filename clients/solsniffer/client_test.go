package solsniffer

import (
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/massigerardi/go-commons/commons"
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
			name: "error test",
			fields: fields{
				apiKey: "bnhd3vwrqaodhmxwn4ry27zwxi9rnl",
				client: httpclient.NewMockRetryableHttpClient([]string{""}, []int{400}),
			},
			args:    args{address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"},
			wantErr: true,
		},
		{
			name: "simple test",
			fields: fields{
				apiKey: "bnhd3vwrqaodhmxwn4ry27zwxi9rnl",
				client: httpclient.NewMockRetryableHttpClient([]string{tokenResponse}, []int{200}),
			},
			args: args{address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN"},
			want: tokenResponse,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expected *TokenData = nil
			if tt.want != "" {
				err := commons.LoadFromJson(tt.want, &expected)
				if err != nil {
					t.Errorf("loading expected data failed: %v", err)
					return
				}
			}
			c := NewClientWithRetryableHttpClient(tt.fields.apiKey, tt.fields.client)
			got, err := c.GetToken(tt.args.address)

			if tt.wantErr && !assert.Error(t, err) {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, expected, got)
		})
	}
}
