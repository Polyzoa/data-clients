package rugcheck

import (
	"time"
)

type TokenData struct {
	Creator    string `json:"creator"`
	DetectedAt string `json:"detectedAt"`
	Events     []struct {
		CreatedAt string `json:"createdAt"`
		Event     int    `json:"event"`
		NewValue  string `json:"newValue"`
		OldValue  string `json:"oldValue"`
	} `json:"events"`
	FileMeta struct {
		Description string `json:"description"`
		Image       string `json:"image"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
	} `json:"fileMeta"`
	FreezeAuthority    string `json:"freezeAuthority"`
	GraphInsiderReport struct {
		Blacklisted  bool `json:"blacklisted"`
		RawGraphData []struct {
			Address string `json:"address"`
			Sent    []struct {
				Amount   int    `json:"amount"`
				Mint     string `json:"mint"`
				Receiver struct {
					Address string `json:"address"`
				} `json:"receiver"`
			} `json:"sent"`
		} `json:"rawGraphData"`
		Receivers []struct {
			Address string `json:"address"`
			Amount  int    `json:"amount"`
		} `json:"receivers"`
		Senders []struct {
			Address string `json:"address"`
			Amount  int    `json:"amount"`
		} `json:"senders"`
		TotalSent int `json:"totalSent"`
	} `json:"graphInsiderReport"`
	KnownAccounts struct {
		AdditionalProp1 struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"additionalProp3"`
	} `json:"knownAccounts"`
	LockerOwners struct {
		AdditionalProp1 bool `json:"additionalProp1"`
		AdditionalProp2 bool `json:"additionalProp2"`
		AdditionalProp3 bool `json:"additionalProp3"`
	} `json:"lockerOwners"`
	Lockers struct {
		AdditionalProp1 struct {
			Owner        string `json:"owner"`
			ProgramID    string `json:"programID"`
			TokenAccount string `json:"tokenAccount"`
			Type         string `json:"type"`
			UnlockDate   int    `json:"unlockDate"`
			Uri          string `json:"uri"`
			UsdcLocked   int    `json:"usdcLocked"`
		} `json:"additionalProp1"`
		AdditionalProp2 struct {
			Owner        string `json:"owner"`
			ProgramID    string `json:"programID"`
			TokenAccount string `json:"tokenAccount"`
			Type         string `json:"type"`
			UnlockDate   int    `json:"unlockDate"`
			Uri          string `json:"uri"`
			UsdcLocked   int    `json:"usdcLocked"`
		} `json:"additionalProp2"`
		AdditionalProp3 struct {
			Owner        string `json:"owner"`
			ProgramID    string `json:"programID"`
			TokenAccount string `json:"tokenAccount"`
			Type         string `json:"type"`
			UnlockDate   int    `json:"unlockDate"`
			Uri          string `json:"uri"`
			UsdcLocked   int    `json:"usdcLocked"`
		} `json:"additionalProp3"`
	} `json:"lockers"`
	LpLockers string `json:"lpLockers"`
	Markets   []struct {
		LiquidityA        string `json:"liquidityA"`
		LiquidityAAccount string `json:"liquidityAAccount"`
		LiquidityB        string `json:"liquidityB"`
		LiquidityBAccount string `json:"liquidityBAccount"`
		Lp                struct {
			Base          int    `json:"base"`
			BaseMint      string `json:"baseMint"`
			BasePrice     int    `json:"basePrice"`
			BaseUSD       int    `json:"baseUSD"`
			CurrentSupply int    `json:"currentSupply"`
			Holders       []struct {
				Address        string `json:"address"`
				Amount         int    `json:"amount"`
				Decimals       int    `json:"decimals"`
				Insider        bool   `json:"insider"`
				Owner          string `json:"owner"`
				Pct            int    `json:"pct"`
				UiAmount       int    `json:"uiAmount"`
				UiAmountString string `json:"uiAmountString"`
			} `json:"holders"`
			LpCurrentSupply     int    `json:"lpCurrentSupply"`
			LpLocked            int    `json:"lpLocked"`
			LpLockedPct         int    `json:"lpLockedPct"`
			LpLockedUSD         int    `json:"lpLockedUSD"`
			LpMaxSupply         int    `json:"lpMaxSupply"`
			LpMint              string `json:"lpMint"`
			LpTotalSupply       int    `json:"lpTotalSupply"`
			LpUnlocked          int    `json:"lpUnlocked"`
			PctReserve          int    `json:"pctReserve"`
			PctSupply           int    `json:"pctSupply"`
			Quote               int    `json:"quote"`
			QuoteMint           string `json:"quoteMint"`
			QuotePrice          int    `json:"quotePrice"`
			QuoteUSD            int    `json:"quoteUSD"`
			ReserveSupply       int    `json:"reserveSupply"`
			TokenSupply         int    `json:"tokenSupply"`
			TotalTokensUnlocked int    `json:"totalTokensUnlocked"`
		} `json:"lp"`
		MarketType    string `json:"marketType"`
		MintA         string `json:"mintA"`
		MintAAccount  string `json:"mintAAccount"`
		MintB         string `json:"mintB"`
		MintBAccount  string `json:"mintBAccount"`
		MintLP        string `json:"mintLP"`
		MintLPAccount string `json:"mintLPAccount"`
		Pubkey        string `json:"pubkey"`
	} `json:"markets"`
	Mint          string `json:"mint"`
	MintAuthority string `json:"mintAuthority"`
	Risks         []struct {
		Description string `json:"description"`
		Level       string `json:"level"`
		Name        string `json:"name"`
		Score       int    `json:"score"`
		Value       string `json:"value"`
	} `json:"risks"`
	Rugged    bool   `json:"rugged"`
	Score     int    `json:"score"`
	Token     string `json:"token"`
	TokenMeta struct {
		Mutable         bool   `json:"mutable"`
		Name            string `json:"name"`
		Symbol          string `json:"symbol"`
		UpdateAuthority string `json:"updateAuthority"`
		Uri             string `json:"uri"`
	} `json:"tokenMeta"`
	TokenProgram    string `json:"tokenProgram"`
	TokenType       string `json:"tokenType"`
	TokenExtensions string `json:"token_extensions"`
	TopHolders      []struct {
		Address        string `json:"address"`
		Amount         int    `json:"amount"`
		Decimals       int    `json:"decimals"`
		Insider        bool   `json:"insider"`
		Owner          string `json:"owner"`
		Pct            int    `json:"pct"`
		UiAmount       int    `json:"uiAmount"`
		UiAmountString string `json:"uiAmountString"`
	} `json:"topHolders"`
	TotalLPProviders     int `json:"totalLPProviders"`
	TotalMarketLiquidity int `json:"totalMarketLiquidity"`
	TransferFee          struct {
		Authority string `json:"authority"`
		MaxAmount int    `json:"maxAmount"`
		Pct       int    `json:"pct"`
	} `json:"transferFee"`
	Verification struct {
		Description string `json:"description"`
		JupVerified bool   `json:"jup_verified"`
		Links       []struct {
			Provider string `json:"provider"`
			Value    string `json:"value"`
		} `json:"links"`
		Mint   string `json:"mint"`
		Name   string `json:"name"`
		Payer  string `json:"payer"`
		Symbol string `json:"symbol"`
	} `json:"verification"`
}

type Token struct {
	Mint            string    `json:"mint"`
	Decimals        int       `json:"decimals"`
	Symbol          string    `json:"symbol"`
	Creator         string    `json:"creator"`
	MintAuthority   string    `json:"mintAuthority"`
	FreezeAuthority string    `json:"freezeAuthority"`
	Program         string    `json:"program"`
	CreateAt        time.Time `json:"createAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Supply          float64   `json:"supply"`
	Events          struct {
		CreatedAt string `json:"createdAt"`
		Event     int    `json:"event"`
		NewValue  string `json:"newValue"`
		OldValue  string `json:"oldValue"`
	} `json:"events"`
}
