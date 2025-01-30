package honeypot

type DataResponse struct {
	Token struct {
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Decimals     int64  `json:"decimals"`
		Address      string `json:"address"`
		TotalHolders int64  `json:"totalHolders"`
	} `json:"token"`
	WithToken struct {
		Name         string `json:"name"`
		Symbol       string `json:"symbol"`
		Decimals     int64  `json:"decimals"`
		Address      string `json:"address"`
		TotalHolders int64  `json:"totalHolders"`
	} `json:"withToken"`
	Summary           Summary `json:"summary"`
	SimulationSuccess bool    `json:"simulationSuccess"`
	HoneypotResult    struct {
		IsHoneypot bool `json:"isHoneypot"`
	} `json:"honeypotResult"`
	SimulationResult struct {
		BuyTax      float64 `json:"buyTax"`
		SellTax     float64 `json:"sellTax"`
		TransferTax float64 `json:"transferTax"`
		BuyGas      float64 `json:"buyGas,string"`
		SellGas     float64 `json:"sellGas,string"`
	} `json:"simulationResult"`
	Flags        []interface{} `json:"flags"`
	ContractCode struct {
		OpenSource     bool `json:"openSource"`
		RootOpenSource bool `json:"rootOpenSource"`
		IsProxy        bool `json:"isProxy"`
		HasProxyCalls  bool `json:"hasProxyCalls"`
	} `json:"contractCode"`
	Chain  Chain  `json:"chain"`
	Router string `json:"router"`
	Pair   struct {
		Pair struct {
			Name    string `json:"name"`
			Address string `json:"address"`
			Token0  string `json:"token0"`
			Token1  string `json:"token1"`
			Type    string `json:"type"`
		} `json:"pair"`
		ChainId            string  `json:"chainId"`
		Reserves0          string  `json:"reserves0"`
		Reserves1          string  `json:"reserves1"`
		Liquidity          float64 `json:"liquidity"`
		Router             string  `json:"router"`
		CreatedAtTimestamp string  `json:"createdAtTimestamp"`
		CreationTxHash     string  `json:"creationTxHash"`
	} `json:"pair"`
	PairAddress string `json:"pairAddress"`
}

type Summary struct {
	Risk      string `json:"risk"`
	RiskLevel int64  `json:"riskLevel"`
	Flags     Flags  `json:"flags"`
}

type Chain struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Currency  string `json:"currency"`
}

type Flags []Flag

type Flag struct {
	Flag          string  `json:"flag"`
	Description   string  `json:"description"`
	Severity      string  `json:"severity"`
	SeverityIndex float64 `json:"severityIndex"`
}
