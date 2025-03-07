package bitquery

import (
	"context"
	"errors"
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
			if got := NewClient(tt.args.apiKey, tt.args.auth); got == nil {
				t.Error("NewClient() is nil")
			}
		})
	}
}

func TestClient_GetStatsData(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		client        *MockGraphqlClient
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
				client:        NewMockClient([]string{TransferResponse}, []error{}),
			},
			args: args{
				chain:     "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    MockTransferData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
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

func checkStats(stats TokenCounters, want TokenCounters, address string, t *testing.T) {
	index := pie.FindFirstUsing(stats, func(value TokenCounter) bool {
		return value.Token.Address == address
	})
	if index == -1 {
		t.Errorf("No data found for %v", address)
		return
	}
	data := stats[index]
	index = pie.FindFirstUsing(want, func(value TokenCounter) bool {
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
		client        *MockGraphqlClient
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
				client:        NewMockClient([]string{ContractResponse}, []error{}),
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
				client:        NewMockClient([]string{}, []error{fmt.Errorf("test error")}),
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
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)

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

func TestClient_GetHoldersData(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		client        *MockGraphqlClient
	}
	type args struct {
		chain   string
		address string
		date    time.Time
		amount  []float64
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
				client:        NewMockClient([]string{MockHoldersResponse}, []error{}),
			},
			args: args{
				chain:   "ethereum",
				address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
				date:    time.Now(),
			},
			want:    HoldersDataResponse,
			wantErr: false,
		},
		{
			name: "Simple Query with amount",
			fields: fields{
				apiKey:        "BQYAZB4gcMCjH5euRUwZQ8WinQWx5UP4",
				authorization: "test",
				client:        NewMockClient([]string{MockHoldersResponse}, []error{}),
			},
			args: args{
				chain:   "ethereum",
				address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
				date:    time.Now(),
				amount:  []float64{1000},
			},
			want:    HoldersDataResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)

			got, err := c.GetHoldersData(context.TODO(), tt.args.chain, tt.args.address, tt.args.date, tt.args.amount...)
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
		client        *MockGraphqlClient
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
				client:        NewMockClient([]string{TransferResponse}, []error{}),
			},
			args: args{
				chainName: "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
				after:     time.Now().AddDate(0, -3, 0),
				before:    time.Now(),
			},
			want:    MockTransferData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)

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

func TestClient_GetHistorySummary(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		client        *MockGraphqlClient
	}
	type args struct {
		chainName string
		address   string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *Counter
		wantErr   bool
		wantCalls int32
	}{
		{
			name: "success",
			fields: fields{
				apiKey:        "test",
				authorization: "test",
				client:        NewMockClient([]string{DataResponse}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			want: &Counter{
				Total:   22223,
				Success: 22167,
				Fails:   56,
			},
			wantErr:   false,
			wantCalls: 1,
		},
		{
			name: "chain error",
			fields: fields{
				apiKey:        "test",
				authorization: "test",
				client:        NewMockClient([]string{}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "etheereum",
			},
			wantErr: true,
		},
		{
			name: "response error",
			fields: fields{
				apiKey:        "test",
				authorization: "test",
				client:        NewMockClient([]string{}, []error{errors.New("response error")}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr:   true,
			wantCalls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)

			got, err := c.GetHistorySummary(tt.args.chainName, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHistorySummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHistorySummary() got = %v, want %v", got, tt.want)
			}
			if tt.fields.client.Calls() != tt.wantCalls {
				t.Errorf("GetHistorySummary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTransfersSummary(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		client        *MockGraphqlClient
	}
	type args struct {
		chainName string
		address   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TransferData
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{TransfersDataResponse}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			want: &TransferData{
				Transfers: []CounterString{
					{
						Total:   835220,
						Success: 821810,
						Fails:   13410,
					},
				},
				Transactions: []CounterString{
					{
						Total:   1498842,
						Success: 1483754,
						Fails:   15088,
					},
				},
			},
		},
		{
			name: "chain failure",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "etheereum",
			},
			wantErr: true,
		},
		{
			name: "resource failure",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{}, []error{errors.New("error")}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			got, err := c.GetTransfersSummary(tt.args.chainName, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransfersSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTransfersSummary()\ngot  %v\nwant %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetHolderBalance(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
		client        *MockGraphqlClient
	}
	type args struct {
		chainName string
		address   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *BalanceData
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{BalanceResponse}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			want: &BalanceData{
				Supply:  142512.74104314,
				Holders: 104594,
			},
		},
		{
			name: "fail empty",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{BalanceResponseEmpty}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr: true,
		},
		{
			name: "fail chain",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{}, []error{}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "etheereum",
			},
			wantErr: true,
		},
		{
			name: "fail response",
			fields: fields{
				apiKey:        "",
				authorization: "",
				client:        NewMockClient([]string{}, []error{errors.New("error")}),
			},
			args: args{
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetClient(tt.fields.client)
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			got, err := c.GetHolderBalance(tt.args.chainName, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHolderBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHolderBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
