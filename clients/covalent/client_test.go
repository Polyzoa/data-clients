package covalent

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/covalenthq/covalent-api-sdk-go/services"
)

func TestClient_GetTransactionSummary(t *testing.T) {
	var count int64 = 1781
	var blockSignedAt, _ = time.Parse("2006-01-02T15:04:05Z", "2024-10-30T17:00:11Z")
	var txHash string = "0x3ea9c546a3143d841b1f880031b7bebfc46bd0ab676a5446c359249c5aa9c6e7"
	type fields struct {
		client CovalentClientType
	}
	type args struct {
		chain   string
		address string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *services.TransactionsSummary
		wantErr bool
	}{
		{
			name:   "test summary",
			fields: fields{client: MockCovalentClient{response: MockSummary}},
			args: args{
				chain:   "ethereum",
				address: "0x2abf0cFB2F36d093F2DBa97bB4214e8FD852ad7c",
			},
			want: &services.TransactionsSummary{
				TotalCount: &count,
				LatestTransaction: &services.TransactionSummary{
					BlockSignedAt: &blockSignedAt,
					TxHash:        &txHash,
				},
				EarliestTransaction: &services.TransactionSummary{
					BlockSignedAt: &blockSignedAt,
					TxHash:        &txHash,
				},
			},
			wantErr: false,
		},
		{
			name:   "test remote error",
			fields: fields{client: MockCovalentClient{err: errors.New("remote")}},
			args: args{
				chain:   "ethereum",
				address: "0x2abf0cFB2F36d093F2DBa97bB4214e8FD852ad7c",
			},
			wantErr: true,
		},
		{
			name:   "test empty error",
			fields: fields{client: NewMockCovalentClient(mockEmptySummary, nil)},
			args: args{
				chain:   "ethereum",
				address: "0x2abf0cFB2F36d093F2DBa97bB4214e8FD852ad7c",
			},
			wantErr: true,
		},
		{
			name:   "test wrong chain",
			fields: fields{client: MockCovalentClient{response: MockSummary}},
			args: args{
				chain:   "ethe",
				address: "0x2abf0cFB2F36d093F2DBa97bB4214e8FD852ad7c",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClientWithMock(tt.fields.client)
			got, err := c.GetTransactionSummary(tt.args.chain, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransactionsSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			jsonData, _ := json.Marshal(got)
			expectedData, _ := json.Marshal(tt.want)
			if !reflect.DeepEqual(string(jsonData), string(expectedData)) {
				t.Errorf("GetTransactionsSummary() \ngot = %v,\n want %v", string(jsonData), string(expectedData))
			}
		})
	}
}
