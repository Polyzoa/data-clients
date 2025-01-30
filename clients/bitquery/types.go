package bitquery

type Response[V any] struct {
	Data   V       `json:"data"`
	Errors []Error `json:"errors"`
}

type Error struct {
	Message   string   `json:"message"`
	Backtrace []string `json:"backtrace"`
	QueryId   string   `json:"query_id"`
}

type StatsData struct {
	TransfersStats []StatData `json:"transfer_stats"`
	Transfers      Counters   `json:"transfers,omitempty"`
	Senders        Counters   `json:"senders,omitempty"`
	Transactions   Counters   `json:"transactions,omitempty"`
}

type Counters []Counter

type Counter struct {
	Token   Token `json:"token,omitempty"`
	Fails   int64 `json:"fails,string,omitempty"`
	Success int64 `json:"success,string,omitempty"`
	Total   int64 `json:"total,string,omitempty"`
}

type Token struct {
	Address  string   `json:"address"`
	Currency Currency `json:"currency"`
}

type Currency struct {
	Address string
}

type ContractData struct {
	Transactions []Contract `json:"transactions"`
}

type Contract struct {
	Hash   string `json:"hash"`
	Sender struct {
		Address       string `json:"address"`
		SmartContract struct {
			ContractType interface{} `json:"contractType"`
		} `json:"smartContract"`
	} `json:"sender"`
	Creates struct {
		Address       string `json:"address"`
		SmartContract struct {
			Currency struct {
				Decimals  int64  `json:"decimals"`
				Name      string `json:"name"`
				Symbol    string `json:"symbol"`
				TokenType string `json:"tokenType"`
			} `json:"currency"`
			ContractType string `json:"contractType"`
		} `json:"smartContract"`
	} `json:"creates"`
	Block struct {
		Timestamp struct {
			Unixtime int64 `json:"unixtime"`
		} `json:"timestamp"`
	} `json:"block"`
}

type Holder struct {
	Holder struct {
		Address string `json:"Address"`
	} `json:"Holder"`
	Balance struct {
		Amount float64 `json:"Amount,string"`
	} `json:"balance"`
	Percentage float64 `json:"percentage,omitempty"`
}

type HoldersData struct {
	BurntBalance []BurntBalance `json:"burnt_balance"`
	HighHolders  []StatData     `json:"high_holders"`
	Holders      []StatData     `json:"holders"`
	LowHolders   []StatData     `json:"low_holders"`
	TopHolders   []Holder       `json:"top_holders"`
}

type StatData struct {
	Token             Token   `json:"token"`
	Receivers         int64   `json:"receivers,string"`
	Total             float64 `json:"total,string"`
	Average           float64 `json:"average"`
	Holders           int64   `json:"holders,string"`
	Median            float64 `json:"median"`
	StandardDeviation float64 `json:"standard_deviation"`
}

type BurntBalance struct {
	Balance float64 `json:"balance"`
}
