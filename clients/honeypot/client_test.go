package honeypot

import (
	"context"
	"reflect"
	"testing"

	"data-clients/clients/internal/httpclient"
)

func TestClient_GetSummary(t *testing.T) {
	type fields struct {
		client httpclient.RetryableHttpClient
	}
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Summary
		wantErr bool
	}{
		{
			name:   "Success call",
			fields: fields{client: httpclient.NewMockRetryableHttpClient([]string{ResponseData}, []int{200})},
			args: args{
				address: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			},
			want: &Summary{
				Risk:      "low",
				RiskLevel: 1,
				Flags:     Flags{},
			},
		},
		// {
		// 	name:   "Success call",
		// 	fields: fields{client: &http.Client{}},
		// 	args: args{
		// 		address: "0x571099c4515345fe09b9ebb0054b2e82970a030b",
		// 	},
		// 	want: &Summary{
		// 		Risk:      "honeypot",
		// 		RiskLevel: 100,
		// 		Flags: Flags{
		// 			{
		// 				Flag:          "high_fail_rate",
		// 				Description:   "A very high amount of users cannot sell their tokens.",
		// 				Severity:      "critical",
		// 				SeverityIndex: 20,
		// 			},
		// 		},
		// 	},
		// },
		{
			name:   "Error call",
			fields: fields{client: httpclient.NewMockRetryableHttpClient([]string{ResponseError}, []int{404})},
			args: args{
				address: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Client{
				client: tt.fields.client,
			}
			got, err := c.GetSummary(context.WithValue(context.TODO(), "id", "test"), tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSummary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
