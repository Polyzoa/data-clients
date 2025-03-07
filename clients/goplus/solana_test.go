package goplus

import (
	"reflect"
	"testing"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/massigerardi/go-commons/commons"
)

func TestSolanaClient_GetSecurity(t *testing.T) {
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		skip    bool
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			skip: true,
			name: "simple live test",
			fields: fields{
				client: retryablehttp.NewClient(),
			},
			args: args{
				address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
			},
			want: MockSolanaSecurityInfo,
		},
		{
			name: "simple test",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{MockSolanaSecurityInfoResult},
					[]int{200},
				),
			},
			args: args{
				address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
			},
			want: MockSolanaSecurityInfo,
		},
		{
			name: "error test",
			fields: fields{
				client: httpclient.NewMockRetryableHttpClient(
					[]string{},
					[]int{400},
				),
			},
			args: args{
				address: "6p6xgHyF7AeE6TZkSmFsko444wqoP15icUSqi2jfGiPN",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				return
			}
			c := SolanaClient{
				client: tt.fields.client,
			}
			got, err := c.GetSecurity(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSecurity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var expected *SolanaSecurityInfo
			if tt.want != "" {
				err = commons.LoadFromJson(tt.want, &expected)
				if err != nil {
					t.Errorf("GetSecurity() json error = %v", err)
					return
				}
			}
			if !reflect.DeepEqual(got, expected) {
				t.Errorf("GetSecurity() \ngot = %v\n want %v", got, expected)
			}
		})
	}
}
