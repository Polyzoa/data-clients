package solsniffer

import (
	"time"
)

type TokenData struct {
	TokenData struct {
		IndicatorData struct {
			High struct {
				Count   int    `json:"count"`
				Details string `json:"details"`
			} `json:"high"`
			Moderate struct {
				Count   int    `json:"count"`
				Details string `json:"details"`
			} `json:"moderate"`
			Low struct {
				Count   int    `json:"count"`
				Details string `json:"details"`
			} `json:"low"`
			Specific struct {
				Count   int    `json:"count"`
				Details string `json:"details"`
			} `json:"specific"`
		} `json:"indicatorData"`
		TokenOverview struct {
			Deployer string `json:"deployer"`
			Mint     string `json:"mint"`
			Address  string `json:"address"`
			Type     string `json:"type"`
		} `json:"tokenOverview"`
		Address       string    `json:"address"`
		TokenName     string    `json:"tokenName"`
		DeployTime    time.Time `json:"deployTime"`
		Externals     string    `json:"externals"`
		LiquidityList []struct {
			Fluxbeam struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
				LpPair  string  `json:"lpPair"`
			} `json:"fluxbeam,omitempty"`
			Raydium struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
				LpPair  string  `json:"lpPair"`
			} `json:"raydium,omitempty"`
			Orca struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
				LpPair  string  `json:"lpPair"`
			} `json:"orca,omitempty"`
			RaydiumCpmm struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
				LpPair  string  `json:"lpPair"`
			} `json:"raydiumCpmm,omitempty"`
			Meteora struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
				LpPair  string  `json:"lpPair"`
			} `json:"meteora,omitempty"`
		} `json:"liquidityList"`
		MarketCap  float64 `json:"marketCap"`
		OwnersList []struct {
			Address    string `json:"address"`
			Amount     string `json:"amount"`
			Percentage string `json:"percentage"`
		} `json:"ownersList"`
		Score       int    `json:"score"`
		TokenImg    string `json:"tokenImg"`
		TokenSymbol string `json:"tokenSymbol"`
		AuditRisk   struct {
			MintDisabled   bool `json:"mintDisabled"`
			FreezeDisabled bool `json:"freezeDisabled"`
			LpBurned       bool `json:"lpBurned"`
			Top10Holders   bool `json:"top10Holders"`
		} `json:"auditRisk"`
	} `json:"tokenData"`
	TokenInfo struct {
		Price        string  `json:"price"`
		SupplyAmount float64 `json:"supplyAmount"`
		MktCap       float64 `json:"mktCap"`
	} `json:"tokenInfo"`
}
