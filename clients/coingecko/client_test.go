package coingecko

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/massigerardi/go-commons/commons"
	"github.com/stretchr/testify/assert"
)

func TestClient_getRemoteCoins(t *testing.T) {
	err := os.Setenv("COINS_FILE", "coins-test.json")
	if err != nil {
		t.Errorf("failed to set env %v", err)
	}
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	tests := []struct {
		skip    bool
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			skip: true,
			name: "Live Query",
			fields: fields{
				client: retryablehttp.NewClient(),
			},
			want:    CoinsResponse,
			wantErr: false,
		},
		{
			name: "Simple Query",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{CoinsResponse}, []int{200}),
			},
			want:    CoinsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if tt.skip {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.getRemoteCoins()
			if (err != nil) != tt.wantErr {
				t.Errorf("getCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			var expected Coins
			_ = commons.LoadFromJson(tt.want, &expected)
			assert.Equal(t, expected, got)
		})
	}
}

func TestClient_GetCoinById(t *testing.T) {
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	type args struct {
		coinId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Simple Query",
			fields: fields{client: httpclient.NewMockRetryableHttpClient([]string{FullCoinResponse}, []int{200})},
			args: args{
				coinId: "bitcoin",
			},
			want:    FullCoinResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.getCoinById(context.WithValue(context.TODO(), "id", "test"), tt.args.coinId)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCoinById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			var expected FullCoin
			_ = commons.LoadFromJson(tt.want, &expected)
			if !reflect.DeepEqual(got, &expected) {
				t.Errorf("getCoinById() \ngot = %v, \n want %v", got, expected)
			}
		})
	}
}

func TestClient_GetCoin(t *testing.T) {
	err := os.Setenv("COINS_FILE", "coins-test.json")
	if err != nil {
		t.Errorf("failed to set env %v", err)
	}
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
			name: "Success Request",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{FullCoinResponse}, []int{
					200,
					200,
				}),
			},
			args: args{address: "0xb8c77482e45f1f44de1745f52c74426c631bdd52", chainName: "ethereum"},
			want: true,
		},
		{
			name:    "Missing Coin List Request",
			fields:  fields{client: httpclient.NewMockRetryableHttpClient([]string{CoinsResponse}, []int{200})},
			args:    args{address: "0xb8c77482e45f1f44de1745f52c74426c631bdd53", chainName: "ethereum"},
			want:    false,
			wantErr: true,
		},
		{
			name: "Missing Coin Request",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient([]string{
					CoinsResponse,
					MissingCoinResponse,
				}, []int{200, 404}),
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
