package goplus

import (
	"testing"

	"github.com/elliotchance/pie/v2"
	"github.com/stretchr/testify/assert"
)

func TestSolanaSecurityInfoFromJson(t *testing.T) {
	type args struct {
		jsonData string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple test",
			args: args{
				jsonData: MockSolanaSecurityInfo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SolanaSecurityInfoFromJson(tt.args.jsonData)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolanaSecurityInfoFromJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, int32(642760), got.HolderCount)
			assert.Equal(t, 10, len(got.Holders))
			holder := pie.First(got.Holders)
			assert.Equal(t, "2RH6rUTPBJ9rUDPpuV9b8z1YL56k1tYU6Uk5ZoaEFFSK", holder.Account)
			assert.Equal(t, 800000022.764, holder.Balance)
			assert.Equal(t, 0, holder.IsLocked)
			assert.Equal(t, 0.8000, holder.Percent)
			assert.Equal(t, "HkykUVWTctptXZmRTWearMsH4SaQNmE4Ku3tMJe5v2mH", holder.TokenAccount)
			assert.Equal(t, 10, len(got.LpHolders))
			holder = pie.First(got.LpHolders)
			assert.Equal(t, "DCn1s6ctS9UhHaxY31XY6pXBMnQv8MQvAy3QAoRohAkS", holder.Account)
			assert.Equal(t, 1036.579805725, holder.Balance)
			assert.Equal(t, 0, holder.IsLocked)
			assert.Equal(t, 0.9622, holder.Percent)
			assert.Equal(t, "4K6YtbpxW2XXhLtVcG3Gm6ZyPNXY7rqtbTJbyKJ4BVAz", holder.TokenAccount)
			assert.Equal(t, "", got.Metadata.Description)
			assert.Equal(t, "OFFICIAL TRUMP", got.Metadata.Name)
			assert.Equal(t, "TRUMP", got.Metadata.Symbol)
			assert.Equal(t, "https://arweave.net/cSCP0h2n1crjeSWE9KF-XtLciJalDNFs7Vf-Sm0NNY0", got.Metadata.Uri)
			assert.Equal(t, 0, got.MetadataMutable.Status)
			assert.Equal(t, 0, got.Mintable.Status)
			assert.Equal(t, 0, got.NonTransferable)
			assert.Equal(t, 999999558.040176, got.TotalSupply)
			assert.Equal(t, 0, got.TrustedToken)
			assert.Equal(t, 0, got.TransferFeeUpgradable.Status)
			assert.Equal(t, 10, len(got.Dex))
			dex := pie.First(got.Dex)
			assert.Equal(t, "raydium", dex.DexName)
			assert.Equal(t, 0.0025, dex.FeeRate)
			assert.Equal(t, 11.209251619252848, dex.Day.PriceMax)
			assert.Equal(t, 12.15741032317363, dex.Week.PriceMax)
			assert.Equal(t, 13.613992254767162, dex.Month.PriceMax)

		})
	}
}

func TestSecurityInfoFromJson(t *testing.T) {
	type args struct {
		jsonData string
	}
	tests := []struct {
		name    string
		args    args
		want    SecurityInfo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple test",
			args: args{
				jsonData: MockSecurityInfo,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SecurityInfoFromJson(tt.args.jsonData)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, float64(0), got.BuyTax)
			assert.Equal(t, 0, got.CanTakeBackOwnership)
			assert.Equal(t, 0, got.CannotBuy)
			assert.Equal(t, 0, got.CannotSellAll)
			assert.Equal(t, 1, got.HiddenOwner)
			assert.Equal(t, int32(40), got.HolderCount)
			assert.Equal(t, 10, len(got.Holders))
			holder := pie.First(got.Holders)
			assert.Equal(t, "0x020bbfdedb16d290dfbc1b6f1cbe08e27a152a6e", holder.Address)
			assert.Equal(t, 1769801017821791.9629249, holder.Balance)
			assert.Equal(t, 0, holder.IsLocked)
			assert.Equal(t, 1769801.017821791962924900, holder.Percent)
			assert.Equal(t, int32(2), got.LpHolderCount)
			assert.Equal(t, 2, len(got.LpHolders))
			holder = pie.First(got.LpHolders)
			assert.Equal(t, "0x000000000000000000000000000000000000dead", holder.Address)
			assert.Equal(t, 0.921954445729287731, holder.Balance)
			assert.Equal(t, 1, holder.IsLocked)
			assert.Equal(t, 0.999999999999998915, holder.Percent)
			assert.Equal(t, 1, len(got.Dex))
			dex := pie.First(got.Dex)
			assert.Equal(t, "UniswapV2", dex.Name)
			assert.Equal(t, 0.00152942, dex.Liquidity)

		})
	}
}
