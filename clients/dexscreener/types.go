package dexscreener

type PairsResponse struct {
	SchemaVersion string `json:"schemaVersion"`
	Pairs         Pairs  `json:"pairs"`
}

type Pairs []Pair

type Pair struct {
	ChainId       string       `json:"chainId"`
	DexId         string       `json:"dexId"`
	Url           string       `json:"url"`
	PairAddress   string       `json:"pairAddress"`
	Labels        []string     `json:"labels"`
	BaseToken     Token        `json:"baseToken"`
	QuoteToken    Token        `json:"quoteToken"`
	PriceNative   float64      `json:"priceNative,string"`
	PriceUsd      float64      `json:"priceUsd,string"`
	Txns          Transactions `json:"txns"`
	Volume        Volume       `json:"volume"`
	PriceChange   PriceChange  `json:"priceChange"`
	Liquidity     Liquidity    `json:"liquidity"`
	Fdv           int          `json:"fdv"`
	MarketCap     int          `json:"marketCap"`
	PairCreatedAt int64        `json:"pairCreatedAt"`
}

type Token struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
}

type Transactions struct {
	M5  BuySell `json:"m5"`
	H1  BuySell `json:"h1"`
	H6  BuySell `json:"h6"`
	H24 BuySell `json:"h24"`
}

type BuySell struct {
	Buys  int `json:"buys"`
	Sells int `json:"sells"`
}

type Volume struct {
	H24 float64 `json:"h24"`
	H6  float64 `json:"h6"`
	H1  float64 `json:"h1"`
	M5  float64 `json:"m5"`
}

type PriceChange struct {
	M5  float64 `json:"m5"`
	H1  float64 `json:"h1"`
	H6  float64 `json:"h6"`
	H24 float64 `json:"h24"`
}

type Liquidity struct {
	Usd   float64 `json:"usd"`
	Base  float64 `json:"base"`
	Quote float64 `json:"quote"`
}

type Profiles []Profile

type Profile struct {
	Url          string  `json:"url"`
	ChainId      string  `json:"chainId"`
	TokenAddress string  `json:"tokenAddress"`
	Amount       float64 `json:"amount,omitempty"`
	TotalAmount  float64 `json:"totalAmount,omitempty"`
	Icon         string  `json:"icon"`
	Header       string  `json:"header,omitempty"`
	OpenGraph    string  `json:"openGraph"`
	Description  string  `json:"description"`
	Links        []struct {
		Label string `json:"label,omitempty"`
		Url   string `json:"url"`
		Type  string `json:"type,omitempty"`
	} `json:"links"`
}
