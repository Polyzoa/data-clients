package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetCoin(t *testing.T) {
	_ = os.Setenv("COINS_FILE", "coins-test.json")
	_ = os.Setenv("ASSETS_FILE", "assets-test.json")
	defer func() {
		_ = os.Remove("coins-test.json")
		_ = os.Remove("assets-test.json")
	}()
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	type args struct {
		chainName string
		address   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "binancecoin",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{
						getMockData("test_coins.json"),
						getMockData("test_assets.json"),
						FullCoinResponse,
					},
					[]int{
						200,
						200,
						200,
					}),
			},
			args: args{address: "0xb8c77482e45f1f44de1745f52c74426c631bdd52", chainName: "ethereum"},
			want: true,
		},
		{
			name: "4Trump",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{
						Coin4Trump,
					},
					[]int{
						200,
						200,
						200,
					}),
			},
			args: args{address: "4TRUMPJwguiFjfY6PLpwT6SZ5BZmCuzfqWWeJAdR6xP3", chainName: "solana"},
			want: true,
		},
		{
			name: "Missing Coin List Request",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{},
					[]int{},
				),
			},
			args:    args{address: "0xb8c77482e45f1f44de1745f52c74426c631bdd53", chainName: "ethereum"},
			want:    false,
			wantErr: true,
		},
		{
			name: "Missing Coin Request",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{
					MissingCoinResponse,
				}, []int{
					404,
				},
				),
			},
			args:    args{address: "0xb8c77482e45f1f44de1745f52c74426c631bdd52", chainName: "ethereum"},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.client)
			got, err := c.GetCoin(context.WithValue(context.TODO(), "id", "test"), tt.args.chainName, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoinByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) != tt.want {
				t.Errorf("GetCoinByToken() want %v", tt.want)
			}
		})
	}
}

func TestClient_getCoins(t *testing.T) {
	localFile := "coins-test.json"
	type fields struct {
		client httpclient.MockRetryableHttpClient
	}
	tests := []struct {
		name      string
		fields    fields
		env       string
		want      string
		wantCalls int32
		wantErr   assert.ErrorAssertionFunc
	}{
		{
			name: "Remote Query",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{MockCoins}, []int{200}),
			},
			env:       localFile,
			want:      MockCoins,
			wantErr:   assert.NoError,
			wantCalls: 1,
		},
		{
			name: "Local Query",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{MockCoins}, []int{200}),
			},
			env:       localFile,
			want:      MockCoins,
			wantErr:   assert.NoError,
			wantCalls: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv("COINS_FILE", tt.env)
			if err != nil {
				t.Errorf("failed to set env %v", err)
			}
			c := &Client{client: tt.fields.client}
			got, err := c.getCoins()
			if !tt.wantErr(t, err, fmt.Sprintf("getCoins()")) {
				return
			}
			var expected Coins
			err = json.Unmarshal([]byte(tt.want), &expected)
			if err != nil {
				t.Error(err)
			}
			assert.Equalf(t, expected, got, "getCoins()")
			assert.Equal(t, tt.wantCalls, tt.fields.client.Calls())
		})
	}
	_ = os.Remove(localFile)
}

func TestClient_getAssets(t *testing.T) {
	type fields struct {
		client httpclient.MockRetryableHttpClient
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		wantErr   assert.ErrorAssertionFunc
		wantCalls int32
	}{
		{
			name: "remote query",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{MockPartialAssets}, []int{200}),
			},
			want:      MockPartialAssets,
			wantErr:   assert.NoError,
			wantCalls: 1,
		},
		{
			name: "local query",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{MockPartialAssets}, []int{200}),
			},
			want:      MockPartialAssets,
			wantErr:   assert.NoError,
			wantCalls: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.Setenv("ASSETS_FILE", "assets-test.json")
			if err != nil {
				t.Errorf("failed to set env %v", err)
			}
			c := &Client{client: tt.fields.client}
			got, err := c.getAssets()
			if !tt.wantErr(t, err, fmt.Sprintf("getCoins()")) {
				return
			}
			var expected Assets
			err = json.Unmarshal([]byte(tt.want), &expected)
			if err != nil {
				t.Error(err)
			}
			assert.Equalf(t, expected, got, "getAssets()")
			assert.Equal(t, tt.wantCalls, tt.fields.client.Calls())
		})
	}
	_ = os.Remove("assets-test.json")
}

func PointerTo[T any](s T) *T {
	return &s
}

func TestClient_GetAsset(t *testing.T) {
	_ = os.Setenv("COINS_FILE", "coins-test.json")
	_ = os.Setenv("ASSETS_FILE", "assets-test.json")
	defer func() {
		_ = os.Remove("coins-test.json")
		_ = os.Remove("assets-test.json")
	}()
	assetsJson := getMockData("test_assets.json")
	var assets Assets
	_ = json.Unmarshal([]byte(assetsJson), &assets)
	type fields struct {
		client httpclient.MockRetryableHttpClient
	}
	type args struct {
		shortname string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "ethereum",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{getMockData("test_coins.json"), getMockData("test_assets.json")},
					[]int{200, 200}),
			},
			args:    args{shortname: "ethereum"},
			want:    0,
			wantErr: assert.NoError,
		},
		{
			name: "bsc",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{getMockData("test_coins.json"), getMockData("test_assets.json")},
					[]int{200, 200}),
			},
			args:    args{shortname: "bsc"},
			want:    352,
			wantErr: assert.NoError,
		},
		{
			name: "solana",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{getMockData("test_coins.json"), getMockData("test_assets.json")},
					[]int{200, 200}),
			},
			args:    args{shortname: "solana"},
			want:    98,
			wantErr: assert.NoError,
		},
		{
			name: "base",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{getMockData("test_coins.json"), getMockData("test_assets.json")},
					[]int{200, 200}),
			},
			args:    args{shortname: "base"},
			want:    363,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.client)
			got, err := c.GetAsset(tt.args.shortname)
			if !tt.wantErr(t, err, fmt.Sprintf("GetAsset(%v)", tt.args.shortname)) {
				return
			}
			expected := assets[tt.want]
			assert.Equalf(t, &expected, got, "GetAsset(%v)", tt.args.shortname)
		})
	}
}
