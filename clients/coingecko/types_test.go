package coingecko

import (
	"testing"
)

func Test_getChain(t *testing.T) {
	type args struct {
		chainName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test ethereum",
			args: args{
				chainName: "ethereum",
			},
			want: true,
		},
		{
			name: "test etheeum",
			args: args{
				chainName: "etheeum",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getChain(tt.args.chainName); (got != nil) != tt.want {
				t.Errorf("getChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
