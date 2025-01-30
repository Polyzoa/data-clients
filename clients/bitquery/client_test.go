package bitquery

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/elliotchance/pie/v2"
	"github.com/massigerardi/go-commons/commons"
	"github.com/massigerardi/graphql"
)

func TestNewClient(t *testing.T) {
	type args struct {
		apiKey   string
		auth     string
		clientV1 graphql.Runner
		clientV2 graphql.Runner
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "new client",
			args: args{
				apiKey:   "test",
				auth:     "test",
				clientV1: graphql.NewClient("https://graphql.bitquery.io"),
				clientV2: graphql.NewClient("https://streaming.bitquery.io/graphql"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiKey, tt.args.auth, tt.args.clientV1, tt.args.clientV2); got == nil {
				t.Error("NewClient() is nil")
			}
		})
	}
}

func TestClient_GetStatsData(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		clientV1      graphql.Runner
		clientV2      graphql.Runner
	}
	type args struct {
		chain     string
		addresses []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple Query",
			fields: fields{
				apiKey:        "test",
				authorization: "test",
				clientV1:      nil,
				clientV2:      NewMockGraphqlClient([]string{TransferResponse}, []error{}),
			},
			args: args{
				chain:     "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    TransferData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				apiKey:        tt.fields.apiKey,
				authorization: tt.fields.authorization,
				clientV1:      tt.fields.clientV1,
				clientV2:      tt.fields.clientV2,
			}
			got, err := c.GetStatsData(context.TODO(), tt.args.chain, tt.args.addresses)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransferData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var expected StatsData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetTransferData() error = %v", err)
			}
			for _, address := range tt.args.addresses {
				checkStats(got.Transfers, expected.Transfers, address, t)
				checkStats(got.Senders, expected.Senders, address, t)
				checkStats(got.Transactions, expected.Transactions, address, t)
			}
		})
	}
}

func checkStats(stats Counters, want Counters, address string, t *testing.T) {
	index := pie.FindFirstUsing(stats, func(value Counter) bool {
		return value.Token.Address == address
	})
	if index == -1 {
		t.Errorf("No data found for %v", address)
	}
	data := stats[index]
	index = pie.FindFirstUsing(want, func(value Counter) bool {
		return value.Token.Address == address
	})
	expected := want[index]
	if data.Total < expected.Total || data.Fails < expected.Fails || data.Success < expected.Success {
		t.Errorf("Wrong data found for %v", address)
	}
}

func TestClient_GetContractData(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		clientV1      graphql.Runner
		clientV2      graphql.Runner
	}
	type args struct {
		chain     string
		addresses []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple Query",
			fields: fields{
				apiKey:        "BQYAZB4gcMCjH5euRUwZQ8WinQWx5UP4",
				authorization: "test",
				clientV1:      NewMockGraphqlClient([]string{ContractResponse}, []error{}),
				clientV2:      nil,
			},
			args: args{
				chain:     "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    ContractDataResponse,
			wantErr: false,
		},
		{
			name: "Error Query",
			fields: fields{
				apiKey:        "BQYAZB4gcMCjH5euRUwZQ8WinQWx5UP4",
				authorization: "test",
				clientV1:      NewMockGraphqlClient([]string{}, []error{fmt.Errorf("test error")}),
				clientV2:      nil,
			},
			args: args{
				chain:     "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				apiKey:        tt.fields.apiKey,
				authorization: tt.fields.authorization,
				clientV1:      tt.fields.clientV1,
				clientV2:      tt.fields.clientV2,
			}
			got, err := c.GetContractData(context.TODO(), tt.args.chain, tt.args.addresses)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContractData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			var expected ContractData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetContractData() error = %v", err)
			}
			if !reflect.DeepEqual(got, &expected) {
				t.Errorf("GetContractData() \ngot  = %v\n want %v", got, expected)
			}
		})
	}
}

func TestClient_getHoldersData(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		clientV1      graphql.Runner
		clientV2      graphql.Runner
	}
	type args struct {
		chain   string
		address string
		date    time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple Query",
			fields: fields{
				apiKey:        "BQYAZB4gcMCjH5euRUwZQ8WinQWx5UP4",
				authorization: "test",
				clientV1:      nil,
				clientV2:      NewMockGraphqlClient([]string{HoldersResponse}, []error{}),
			},
			args: args{
				chain:   "ethereum",
				address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
				date:    time.Now(),
			},
			want:    HoldersDataResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				apiKey:        tt.fields.apiKey,
				authorization: tt.fields.authorization,
				clientV1:      tt.fields.clientV1,
				clientV2:      tt.fields.clientV2,
			}
			got, err := c.GetHoldersData(context.TODO(), tt.args.chain, tt.args.address, tt.args.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHoldersData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var expected HoldersData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetContractData() error = %v", err)
			}
			if !reflect.DeepEqual(got, &expected) {
				t.Errorf("GetContractData() \ngot  = %v\n want %v", got, expected)
			}
		})
	}
}

func TestClient_GetStatsDataBefore(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		clientV1      graphql.Runner
		clientV2      graphql.Runner
	}
	type args struct {
		in0       context.Context
		chainName string
		addresses []string
		after     time.Time
		before    time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Simple Query",
			fields: fields{
				apiKey:        "test",
				authorization: "test",
				clientV1:      nil,
				clientV2:      NewMockGraphqlClient([]string{TransferResponse}, []error{}),
			},
			args: args{
				chainName: "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
				after:     time.Now().AddDate(0, -3, 0),
				before:    time.Now(),
			},
			want:    TransferData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				apiKey:        tt.fields.apiKey,
				authorization: tt.fields.authorization,
				clientV1:      tt.fields.clientV1,
				clientV2:      tt.fields.clientV2,
			}
			got, err := c.GetStatsDataBefore(tt.args.in0, tt.args.chainName, tt.args.addresses, tt.args.after, tt.args.before)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatsDataBefore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var expected StatsData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetTransferData() error = %v", err)
			}
			for _, address := range tt.args.addresses {
				checkStats(got.Transfers, expected.Transfers, address, t)
				checkStats(got.Senders, expected.Senders, address, t)
				checkStats(got.Transactions, expected.Transactions, address, t)
			}
		})
	}
}
