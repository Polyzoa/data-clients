package eth

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestClient_IsToken(t *testing.T) {
	type args struct {
		chain   string
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "BNB",
			args: args{
				chain:   "ethereum",
				address: "0xB8c77482e45F1F44dE1745F52C74426C631bDD52",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{}
			client, err := ethclient.Dial(urls[tt.args.chain])
			if err != nil {
				t.Errorf("failed eth client %v", err)
			}
			got, err := c.IsToken(tt.args.address, client)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTokens(t *testing.T) {
	type args struct {
		chain     string
		addresses []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]bool
		wantErr bool
	}{
		{
			name: "BNB",
			args: args{
				chain:     "ethereum",
				addresses: []string{"0xB8c77482e45F1F44dE1745F52C74426C631bDD52", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", "0x558FA75074cc7cF045C764aEd47D37776Ea697d2"},
			},
			want:    map[string]bool{"0xB8c77482e45F1F44dE1745F52C74426C631bDD52": true, "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48": true, "0x558FA75074cc7cF045C764aEd47D37776Ea697d2": false},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{}
			got, err := c.GetTokens(tt.args.chain, tt.args.addresses)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTokens() got = %v, want %v", got, tt.want)
			}
		})
	}
}
