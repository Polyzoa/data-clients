package dexscreener

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/massigerardi/go-commons/commons"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetPairs(t *testing.T) {
	type fields struct {
		client httpclient.HttpClient
	}
	type args struct {
		addresses []string
	}
	tests := []struct {
		fields  fields
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{PairsDataResponse}, []int{200}),
			},
			args: args{
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    PairsData,
			wantErr: false,
		},
		{
			name: "Error Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{PairsDataResponse}, []int{404}),
			},
			args: args{
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    PairsData,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{client: tt.fields.client}
			got, err := c.GetPairs(tt.args.addresses)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPairs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			var expected Pairs
			_ = commons.LoadFromJson(tt.want, &expected)
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("GetPairs() \ngot = %v, \nwant  %v", got, expected)
			}
		})
	}
}

func TestClient_GetLatestTokens(t *testing.T) {
	type fields struct {
		client httpclient.HttpClient
	}
	type args struct {
		network []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          string
		wantErr       assert.ErrorAssertionFunc
		expectedError error
	}{
		{
			name: "Simple Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			want: ProfilesDataResponse,
			args: args{
				network: []string{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Error Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{400}),
			},
			want: "[]",
			args: args{
				network: []string{},
			},
			wantErr:       assert.Error,
			expectedError: fmt.Errorf(http.StatusText(http.StatusBadRequest)),
		},
		{
			name: "Query with Solana",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			want: ProfilesDataResponseSolana,
			args: args{
				network: []string{"solana"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Query with Base",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			want: ProfilesDataResponseBase,
			args: args{
				network: []string{"base"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Query with Celo",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			want: "[]",
			args: args{
				network: []string{"celo"},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.GetLatestTokens(tt.args.network...)
			if tt.wantErr(t, err) {
				assert.Equal(t, tt.expectedError, err)
				return
			}
			var expected Profiles
			err = commons.LoadFromJson(tt.want, &expected)
			assert.Equal(t, expected, got)
		})
	}
}

func TestClient_GetLatestBoostedTokens(t *testing.T) {
	type fields struct {
		client httpclient.HttpClient
	}
	type args struct {
		networks []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          string
		wantErr       assert.ErrorAssertionFunc
		expectedError error
	}{
		{
			name: "Simple Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			args:    args{networks: []string{}},
			want:    ProfilesDataResponse,
			wantErr: assert.NoError,
		},
		{
			name: "Error Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataBoostedResponse}, []int{400}),
			},
			want: "[]",
			args: args{
				networks: []string{},
			},
			wantErr:       assert.Error,
			expectedError: fmt.Errorf(http.StatusText(http.StatusBadRequest)),
		},
		{
			name: "Simple Query 2",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataBoostedResponse}, []int{200}),
			},
			args:    args{networks: []string{}},
			want:    ProfilesDataBoostedResponse,
			wantErr: assert.NoError,
		},
		{
			name: "Query with Solana",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataBoostedResponse}, []int{200}),
			},
			want: ProfilesDataBoostedResponse,
			args: args{
				networks: []string{"solana"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Query with Celo",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			want: "[]",
			args: args{
				networks: []string{"celo"},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.GetLatestBoostedTokens(tt.args.networks...)
			if tt.wantErr(t, err) {
				assert.Equal(t, tt.expectedError, err)
				return
			}
			var expected Profiles
			err = commons.LoadFromJson(tt.want, &expected)
			assert.Equal(t, expected, got)
		})
	}
}
func TestClient_GetTopBoostedTokens(t *testing.T) {
	type fields struct {
		client httpclient.HttpClient
	}
	type args struct {
		networks []string
	}
	tests := []struct {
		name          string
		fields        fields
		args          args
		want          string
		wantErr       assert.ErrorAssertionFunc
		expectedError error
	}{
		{
			name: "Simple Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			args:    args{networks: []string{}},
			want:    ProfilesDataResponse,
			wantErr: assert.NoError,
		},
		{
			name: "Error Query",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataBoostedResponse}, []int{400}),
			},
			want: "[]",
			args: args{
				networks: []string{},
			},
			wantErr:       assert.Error,
			expectedError: fmt.Errorf(http.StatusText(http.StatusBadRequest)),
		},
		{
			name: "Simple Query 2",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataBoostedResponse}, []int{200}),
			},
			args:    args{networks: []string{}},
			want:    ProfilesDataBoostedResponse,
			wantErr: assert.NoError,
		},
		{
			name: "Query with Solana",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataBoostedResponse}, []int{200}),
			},
			want: ProfilesDataBoostedResponse,
			args: args{
				networks: []string{"solana"},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Query with Celo",
			fields: fields{
				client: httpclient.NewMockHttpClient([]string{ProfilesDataResponse}, []int{200}),
			},
			want: "[]",
			args: args{
				networks: []string{"celo"},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.GetTopBoostedTokens(tt.args.networks...)
			if tt.wantErr(t, err) {
				assert.Equal(t, tt.expectedError, err)
				return
			}
			var expected Profiles
			err = commons.LoadFromJson(tt.want, &expected)
			assert.Equal(t, expected, got)
		})
	}
}
