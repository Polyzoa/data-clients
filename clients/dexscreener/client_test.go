package dexscreener

import (
	"reflect"
	"testing"

	"data-clients/clients/internal/httpclient"
	"github.com/massigerardi/go-commons/commons"
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
