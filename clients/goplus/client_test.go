package goplus

import (
	"context"
	"reflect"
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/massigerardi/go-commons/commons"
)

func TestGetTokenInfo(t *testing.T) {

	type fields struct {
		client httpclient.RetryableHttpClient
	}
	type args struct {
		chain   string
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test success token",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{MockSecurityInfoResults},
					[]int{200},
				),
			},
			args: args{
				chain:   "ethereum",
				address: "0x4e7e15a6b83718ba79674cc4ce09f7d4a631ec2c",
			},
			want: MockSecurityInfo,
		},
		// {
		// 	name: "test fail token",
		// 	fields: fields{
		// 		token: MockTokenClient{
		// 			Response: TokenResponse,
		// 			Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		// 			Code:     errorcode.INVALID_TOKEN,
		// 			Message:  "Invalid Token",
		// 		},
		// 		address: nil,
		// 	},
		// 	args: args{
		// 		chain:   "1",
		// 		address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		// 	},
		// 	wantErr: true,
		// },
		// {
		// 	name: "test chain not supported",
		// 	fields: fields{
		// 		token: MockTokenClient{
		// 			Response: TokenResponse,
		// 			Address:  "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		// 			Code:     errorcode.INVALID_TOKEN,
		// 			Message:  "Invalid Token",
		// 		},
		// 		address: nil,
		// 	},
		// 	args: args{
		// 		chain:   "",
		// 		address: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		// 	},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(tt.fields.client)
			got, err := client.GetTokenInfo(context.WithValue(context.TODO(), "id", "test"), tt.args.chain, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTokenInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var expected *SecurityInfo
			err = commons.LoadFromJson(tt.want, &expected)
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("Expected %v\nReceived %v", expected, got)
			}
		})
	}
}
