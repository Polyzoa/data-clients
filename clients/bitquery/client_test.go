package bitquery

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/elliotchance/pie/v2"
	"github.com/massigerardi/go-commons/commons"
	"github.com/massigerardi/graphql"
	"github.com/stretchr/testify/assert"
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
	}
	type args struct {
		chain     string
		addresses []string
		client    GraphqlClient
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
			},
			args: args{
				client: NewMockClient([]string{TransferResponse}, []error{}),
				chain:  "ethereum",
				addresses: []string{
					"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					"0xdac17f958d2ee523a2206206994597c13d831ec7",
				},
			},
			want:    MockTransferData,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			var response Response[StatsData]
			query := Query{
				Client: tt.args.client,
				Query:  StatsQuery.Query,
				Params: map[string]any{
					"addresses": tt.args.addresses,
					"network":   tt.args.chain,
				},
			}
			err := c.RunQuery(&query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransferData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := response.Data
			var expected StatsData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetTransferData() error = %v", err)
			}
			addresses := tt.args.addresses
			for _, address := range addresses {
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
	}
	type args struct {
		chain     string
		addresses []string
		client    GraphqlClient
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
			},
			args: args{
				client:    NewMockClient([]string{ContractResponse}, []error{}),
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
			},
			args: args{
				client:    NewMockClient([]string{}, []error{fmt.Errorf("test error")}),
				chain:     "ethereum",
				addresses: []string{"0xa0b86991c6218b36c1d19d4a9eb0ce3606eb48", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			var response Response[ContractData]
			query := Query{
				Client: tt.args.client,
				Query:  StatsQuery.Query,
				Params: map[string]any{
					"addresses": tt.args.addresses,
					"network":   tt.args.chain,
				},
			}
			err := c.RunQuery(&query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContractData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			got := response.Data
			var expected ContractData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetContractData() error = %v", err)
			}
			assert.Equal(t, expected, got)
		})
	}
}

func TestClient_GetHoldersData(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
	}
	type args struct {
		client  GraphqlClient
		chain   string
		address string
		date    time.Time
		amount  float64
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
			},
			args: args{
				client:  NewMockClient([]string{MockHoldersResponse}, []error{}),
				chain:   "ethereum",
				address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
				date:    time.Now(),
				amount:  1000,
			},
			want:    HoldersDataResponse,
			wantErr: false,
		},
		{
			name: "Simple Query with amount",
			fields: fields{
				apiKey:        "BQYAZB4gcMCjH5euRUwZQ8WinQWx5UP4",
				authorization: "test",
			},
			args: args{
				client:  NewMockClient([]string{MockHoldersResponse}, []error{}),
				chain:   "ethereum",
				address: "0xdac17f958d2ee523a2206206994597c13d831ec7",
				date:    time.Now(),
				amount:  1000,
			},
			want:    HoldersDataResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)

			var response Response[HoldersData]
			query := Query{
				Client: tt.args.client,
				Query:  HoldersQuery.Query,
				Params: map[string]any{
					"address": tt.args.address,
					"network": tt.args.chain,
					"date":    tt.args.date.Format(time.DateOnly),
					"amount":  fmt.Sprintf("%f", tt.args.amount),
				},
			}
			err := c.RunQuery(&query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("getHoldersData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := response.Data
			var expected HoldersData
			err = commons.LoadFromJson(tt.want, &expected)
			if err != nil {
				t.Errorf("GetContractData() error = %v", err)
			}
			assert.Equal(t, expected, got)
		})
	}
}

func TestClient_GetStatsDataBefore(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
	}
	type args struct {
		client    GraphqlClient
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
			},
			args: args{
				client:    NewMockClient([]string{TransferResponse}, []error{}),
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
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			query := StatsQueryInTime.Clone()
			query.Params = map[string]any{
				"addresses": tt.args.addresses,
				"network":   tt.args.chainName,
				"after":     tt.args.after.Format("2006-01-02T15:04:05Z"),
				"before":    tt.args.before.Format("2006-01-02T15:04:05Z"),
			}
			query.Client = tt.args.client
			var response Response[StatsData]
			err := c.RunQuery(query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStatsDataBefore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := response.Data
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
	}
	type args struct {
		client    *MockGraphqlClient
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
			},
			args: args{
				client:    NewMockClient([]string{DataResponse}, []error{}),
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
			name: "response error",
			fields: fields{
				apiKey:        "test",
				authorization: "test",
			},
			args: args{
				client:    NewMockClient([]string{}, []error{errors.New("response error")}),
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr:   true,
			wantCalls: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			query := Query{
				Client: tt.args.client,
				Query:  SuccessTransactionsQuery.Query,
				Params: map[string]any{
					"address": tt.args.address,
					"chain":   tt.args.chainName,
				},
			}

			var response Response[TransactionData]
			err := c.RunQuery(&query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHistorySummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			got := pie.First(response.Data.Results)
			assert.Equal(t, tt.want, got)
			if tt.args.client.Calls() != tt.wantCalls {
				t.Errorf("GetHistorySummary() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetTransfersSummary(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
	}
	type args struct {
		client    *MockGraphqlClient
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
			},
			args: args{
				client:    NewMockClient([]string{TransfersDataResponse}, []error{}),
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
			name: "resource failure",
			fields: fields{
				apiKey:        "",
				authorization: "",
			},
			args: args{
				client:    NewMockClient([]string{}, []error{errors.New("error")}),
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			query := Query{
				Client: tt.args.client,
				Query:  QueryTransactionStats.Query,
				Params: map[string]any{
					"address": tt.args.address,
					"chain":   tt.args.chainName,
				},
			}
			var response Response[TransferData]
			err := c.RunQuery(&query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransfersSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			got := response.Data
			assert.Equal(t, tt.want, &got)
		})
	}
}

func TestClient_GetHolderBalance(t *testing.T) {
	type fields struct {
		apiKey        string
		authorization string
	}
	type args struct {
		client    *MockGraphqlClient
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
			},
			args: args{
				client:    NewMockClient([]string{BalanceResponse}, []error{}),
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			want: &BalanceData{
				Supply:  142512.74104314,
				Holders: 104594,
			},
		},
		{
			name: "fail response",
			fields: fields{
				apiKey:        "",
				authorization: "",
			},
			args: args{
				client:    NewMockClient([]string{}, []error{errors.New("error")}),
				address:   "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
				chainName: "ethereum",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(tt.fields.apiKey, tt.fields.authorization)
			query := Query{
				Client: tt.args.client,
				Query:  BalanceQuery.Query,
				Params: map[string]any{
					"address": tt.args.address,
					"chain":   tt.args.chainName,
				},
			}
			var response Response[BalanceDataResponse]
			err := c.RunQuery(&query, &response)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHolderBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			got := pie.First(response.Data.Holders)
			assert.Equal(t, tt.want, &got)
		})
	}
}
