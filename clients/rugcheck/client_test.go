package rugcheck

import (
	"fmt"
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/massigerardi/go-commons/commons"
	"github.com/stretchr/testify/assert"
)

func getExpected[T any](jsonData string, expected T) T {
	err := commons.LoadFromJson(jsonData, &expected)
	if err != nil {
		fmt.Printf("loading expected data failed: %v", err)
		panic(err)
	}
	return expected
}

func TestClient_GetTokenSummary(t *testing.T) {
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TokenData
		wantErr bool
	}{
		{
			name: "success test",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{tokenResponse}, []int{200}),
			},
			args: args{
				address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
			},
			want: getExpected(tokenResponse, &TokenData{}),
		},
		{
			name: "error test",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{""}, []int{400}),
			},
			args: args{
				address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithRetryableHttpClient(tt.fields.client)
			got, err := c.GetTokenSummary(tt.args.address)
			if tt.wantErr && !assert.Error(t, err) {
				t.Errorf("GetTokenSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetLatestTokens(t *testing.T) {
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*Token
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success test",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{newTokens}, []int{200}),
			},
			want:    getExpected(newTokens, []*Token{}),
			wantErr: assert.NoError,
		},
		{
			name: "fail test",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{""}, []int{400}),
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithRetryableHttpClient(tt.fields.client)
			got, err := c.GetLatestTokens()
			if !tt.wantErr(t, err, fmt.Sprintf("GetLatestTokens()")) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetLatestTokens()")
		})
	}
}
