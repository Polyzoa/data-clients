package goplus

import (
	"encoding/json"
)

type Payload[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}

type SecurityInfoResult map[string]SecurityInfo

type SecurityInfo struct {
	AntiWhaleModifiable        string   `json:"anti_whale_modifiable"`
	BuyTax                     float64  `json:"buy_tax,string"`
	CanTakeBackOwnership       int      `json:"can_take_back_ownership,string"`
	CannotBuy                  int      `json:"cannot_buy,string"`
	CannotSellAll              int      `json:"cannot_sell_all,string"`
	CreatorAddress             string   `json:"creator_address"`
	CreatorBalance             float64  `json:"creator_balance,string"`
	CreatorPercent             float64  `json:"creator_percent,string"`
	Dex                        []Dex    `json:"dex"`
	ExternalCall               string   `json:"external_call"`
	HiddenOwner                int      `json:"hidden_owner,string"`
	HolderCount                int32    `json:"holder_count,string,omitempty"`
	Holders                    []Holder `json:"holders"`
	HoneypotWithSameCreator    int      `json:"honeypot_with_same_creator,string"`
	IsAntiWhale                int      `json:"is_anti_whale,string"`
	IsBlacklisted              int      `json:"is_blacklisted,string"`
	IsHoneypot                 int      `json:"is_honeypot,string"`
	IsInDex                    int      `json:"is_in_dex,string"`
	IsMintable                 int      `json:"is_mintable,string"`
	IsOpenSource               int      `json:"is_open_source,string"`
	IsProxy                    int      `json:"is_proxy,string"`
	IsWhitelisted              int      `json:"is_whitelisted,string"`
	LpHolderCount              int32    `json:"lp_holder_count,string,omitempty"`
	LpHolders                  []Holder `json:"lp_holders"`
	LpTotalSupply              float64  `json:"lp_total_supply,string"`
	OwnerAddress               string   `json:"owner_address"`
	OwnerBalance               float64  `json:"owner_balance,string"`
	OwnerChangeBalance         int      `json:"owner_change_balance,string"`
	OwnerPercent               float64  `json:"owner_percent,string"`
	PersonalSlippageModifiable string   `json:"personal_slippage_modifiable"`
	Selfdestruct               string   `json:"selfdestruct"`
	SellTax                    float64  `json:"sell_tax,string"`
	SlippageModifiable         string   `json:"slippage_modifiable"`
	TokenName                  string   `json:"token_name"`
	TokenSymbol                string   `json:"token_symbol"`
	TotalSupply                float64  `json:"total_supply,string"`
	TradingCooldown            string   `json:"trading_cooldown"`
	TransferPausable           string   `json:"transfer_pausable"`
	TrustList                  int      `json:"trust_list,string,omitempty"`
}

type Holder struct {
	Address      string        `json:"address,omitempty"`
	Account      string        `json:"account,omitempty"`
	Tag          string        `json:"tag,omitempty"`
	IsContract   int           `json:"is_contract,omitempty"`
	Balance      float64       `json:"balance,string,omitempty"`
	Percent      float64       `json:"percent,string,omitempty"`
	IsLocked     int           `json:"is_locked,omitempty"`
	LockedDetail []interface{} `json:"locked_detail,omitempty"`
	TokenAccount string        `json:"token_account,omitempty"`
	Value        interface{}   `json:"value"`
	NFTList      interface{}   `json:"NFT_list"`
}

type Dex struct {
	LiquidityType string  `json:"liquidity_type"`
	Name          string  `json:"name"`
	Liquidity     float64 `json:"liquidity,string"`
	Pair          string  `json:"pair"`
}

type SolanaSecurityInfoResult map[string]SolanaSecurityInfo

type DexStats struct {
	PriceMax float64 `json:"price_max,string"`
	PriceMin float64 `json:"price_min,string"`
	Volume   float64 `json:"volume,string"`
}

type SolanaSecurityInfo struct {
	BalanceMutableAuthority struct {
		Authority []interface{} `json:"authority"`
		Status    string        `json:"status"`
	} `json:"balance_mutable_authority"`
	Closable struct {
		Authority []interface{} `json:"authority"`
		Status    int           `json:"status,string"`
	} `json:"closable"`
	Creators []struct {
		Address          string `json:"address"`
		MaliciousAddress int    `json:"malicious_address"`
	} `json:"creators"`
	DefaultAccountState           string `json:"default_account_state"`
	DefaultAccountStateUpgradable struct {
		Authority []interface{} `json:"authority"`
		Status    string        `json:"status"`
	} `json:"default_account_state_upgradable"`
	Dex []struct {
		Day      DexStats `json:"day"`
		Week     DexStats `json:"week"`
		Month    DexStats `json:"month"`
		DexName  string   `json:"dex_name"`
		FeeRate  float64  `json:"fee_rate,string"`
		Id       string   `json:"id"`
		LpAmount string   `json:"lp_amount"`
		OpenTime string   `json:"open_time"`
		Price    float64  `json:"price,string"`
		Tvl      float64  `json:"tvl,string"`
		Type     string   `json:"type"`
	} `json:"dex"`
	Freezable struct {
		Authority []interface{} `json:"authority"`
		Status    int           `json:"status,string"`
	} `json:"freezable"`
	HolderCount int32    `json:"holder_count,string"`
	Holders     []Holder `json:"holders"`
	LpHolders   []Holder `json:"lp_holders"`
	Metadata    struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		Uri         string `json:"uri"`
	} `json:"metadata"`
	MetadataMutable struct {
		MetadataUpgradeAuthority []interface{} `json:"metadata_upgrade_authority"`
		Status                   int           `json:"status,string"`
	} `json:"metadata_mutable"`
	Mintable struct {
		Authority []interface{} `json:"authority"`
		Status    int           `json:"status,string"`
	} `json:"mintable"`
	NonTransferable int     `json:"non_transferable,string"`
	TotalSupply     float64 `json:"total_supply,string"`
	TransferFee     struct {
		CurrentFeeRate float64 `json:"current_fee_rate,string"`
		FeeRate        float64 `json:"fee_rate,string"`
		MaximumFee     float64 `json:"maximum_fee,string"`
	} `json:"transfer_fee"`
	TransferFeeUpgradable struct {
		Authority []interface{} `json:"authority"`
		Status    int           `json:"status,string"`
	} `json:"transfer_fee_upgradable"`
	TransferHook           []interface{} `json:"transfer_hook"`
	TransferHookUpgradable struct {
		Authority []interface{} `json:"authority"`
		Status    string        `json:"status"`
	} `json:"transfer_hook_upgradable"`
	TrustedToken int `json:"trusted_token"`
}

func SolanaSecurityInfoFromJson(jsonData string) (SolanaSecurityInfo, error) {
	var info SolanaSecurityInfo
	err := json.Unmarshal([]byte(jsonData), &info)
	return info, err
}

func SecurityInfoFromJson(jsonData string) (SecurityInfo, error) {
	var info SecurityInfo
	err := json.Unmarshal([]byte(jsonData), &info)
	return info, err
}
