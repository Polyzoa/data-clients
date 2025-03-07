//go:build live

package bitquery

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_RunQuery(t *testing.T) {
	t.Skip("live test")
	type fields struct {
		apiKey        string
		authorization string
	}
	type params map[string]any
	type args struct {
		query Query
	}
	type test[V any] struct {
		name    string
		fields  fields
		args    args
		want    Response[V]
		wantErr bool
		skip    bool
	}
	tests := []test[any]{
		{
			name: "test Transfers",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: Query{
					Url:   EndpointV2,
					Query: SuccessTransfersQueryV2.Query,
					Params: params{
						"address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
						"chain":   "ethereum",
					},
				},
			},
			want: Response[any]{},
			skip: true,
		},
		{
			skip: true,
			name: "test no params",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: SuccessTransfersQueryV2,
			},
			wantErr: true,
		},
		{
			skip: true,
			name: "test Balance",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: Query{
					Url:   EndpointV2,
					Query: BalanceQuery.Query,
					Params: params{
						"address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
						"chain":   "ethereum",
						"date":    time.Now().Format(time.DateOnly),
					},
				},
			},
			want: Response[any]{},
		},
		{
			skip: true,
			name: "test StatsQuery",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: Query{
					Url:   StatsQuery.Url,
					Query: StatsQuery.Query,
					Params: params{
						"chain": "ethereum",
						"addresses": []string{
							"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
							"0xdac17f958d2ee523a2206206994597c13d831ec7",
						},
					},
					Response: StatsQuery.Response,
				},
			},
			want: Response[any]{},
		},
		{
			skip: true,
			name: "test Stats",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: Query{
					Url:   EndpointV2,
					Query: QueryTransactionStats.Query,
					Params: params{
						"address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
						"chain":   "ethereum",
						"from":    time.Now().Format(time.DateOnly),
						"till":    time.Now().AddDate(0, -3, 0).Format(time.DateOnly),
					},
				},
			},
			want: Response[any]{},
		},
		{
			name: "test Solana",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: Query{
					Url:   SolanaTransferQuery.Url,
					Query: SolanaTransferQuery.Query,
					Params: params{
						"address": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
					},
				},
			},
			want: Response[any]{},
		},
		{
			name: "test Solana",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
			},
			args: args{
				query: Query{
					Url:   SolanaTransferQuery.Url,
					Query: SolanaTransferQuery.Query,
					Params: params{
						"address": "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
					},
				},
			},
			want: Response[any]{},
		},
	}
	for _, tt := range tests {
		if tt.skip {
			continue
		}
		service := NewClient(tt.fields.apiKey, tt.fields.authorization)
		err := service.RunQuery(&tt.args.query)
		if (err != nil) != tt.wantErr {
			t.Errorf("RunQuery() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		println(fmt.Sprintf("results: %v", tt.want))
	}

}
