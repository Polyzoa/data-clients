package test

import (
	"reflect"
	"testing"

	"github.com/Polyzoa/data-clients/clients/bitquery"
)

func TestGetChainV2(t *testing.T) {
	type args struct {
		chainName string
	}
	tests := []struct {
		name    string
		args    args
		want    *bitquery.Chain
		wantErr bool
	}{
		{
			name: "Find Ethereum",
			args: args{chainName: "ethereum"},
			want: &bitquery.Chain{Name: "ethereum", Network: "eth"},
		},
		{
			name:    "Find nil",
			args:    args{chainName: "etheeum"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bitquery.GetChainV2(tt.args.chainName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChainV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChainV2() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetChainV1(t *testing.T) {
	type args struct {
		chainName string
	}
	tests := []struct {
		name    string
		args    args
		want    *bitquery.Chain
		wantErr bool
	}{
		{
			name: "Find Ethereum",
			args: args{chainName: "ethereum"},
			want: &bitquery.Chain{Name: "ethereum", Network: "ethereum"},
		},
		{
			name:    "Find nil",
			args:    args{chainName: "etheeum"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bitquery.GetChainV1(tt.args.chainName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChainV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetChainV2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
