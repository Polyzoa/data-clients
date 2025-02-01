package bitquery

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/massigerardi/graphql"
)

func Test_RunQuery(t *testing.T) {
	// t.Skip("live test")
	// clientV1 := graphql.NewClient("https://graphql.bitquery.io")
	clientV2 := graphql.NewClient("https://streaming.bitquery.io/graphql")
	type fields struct {
		apiKey        string
		authorization string
		clientV1      graphql.Runner
		clientV2      graphql.Runner
	}
	type params map[string]any
	type args struct {
		query  string
		params params
		client graphql.Runner
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
				clientV1:      nil,
				clientV2:      nil,
			},
			args: args{
				client: clientV2,
				query:  SuccessTransfersQueryV2,
				params: params{
					"address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
					"chain":   "ethereum",
				},
			},
			want: Response[any]{},
			skip: true,
		},
		{
			name: "test Balance",
			fields: fields{
				apiKey:        os.Getenv("API_KEY"),
				authorization: "",
				clientV1:      nil,
				clientV2:      nil,
			},
			args: args{
				client: clientV2,
				query:  BalanceQuery,
				params: params{
					"address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
					"chain":   "ethereum",
					"date":    time.Now().Format(time.DateOnly),
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
				clientV1:      nil,
				clientV2:      nil,
			},
			args: args{
				client: clientV2,
				query:  QueryTransactionStats,
				params: params{
					"address": "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
					"chain":   "ethereum",
					"from":    time.Now().Format(time.DateOnly),
					"till":    time.Now().AddDate(0, -3, 0).Format(time.DateOnly),
				},
			},
			want: Response[any]{},
		},
	}
	for _, tt := range tests {
		if tt.skip {
			continue
		}
		service := NewClient(tt.fields.apiKey, tt.fields.authorization, tt.fields.clientV1, tt.fields.clientV2)
		err := service.runQuery(tt.args.client, tt.args.query, tt.args.params, &tt.want)
		if (err != nil) != tt.wantErr {
			t.Errorf("runQuery() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		println(fmt.Sprintf("results: %v", tt.want))
	}

}
