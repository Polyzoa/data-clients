package bitquery

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/massigerardi/graphql"
)

type Endpoint string

const (
	EndpointV1 Endpoint = "https://graphql.bitquery.io"
	EndpointV2 Endpoint = "https://streaming.bitquery.io/graphql"
	EndpointV3 Endpoint = "https://streaming.bitquery.io/eap"
)

type Query struct {
	Client GraphqlClient
	Query  string
	Params map[string]any
}

func (q Query) Clone() *Query {
	buf := bytes.Buffer{}
	err := gob.NewEncoder(&buf).Encode(q)
	if err != nil {
		return &q
	}
	var query Query
	err = gob.NewDecoder(&buf).Decode(&query)
	if err != nil {
		return &q
	}
	return &query
}

func (q Query) Run(ctx context.Context, req *graphql.Request, response interface{}) error {
	return q.Client.Run(ctx, req, response)
}

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
	TransfersStats []StatData    `json:"transfer_stats"`
	Transfers      TokenCounters `json:"transfers,omitempty"`
	Senders        TokenCounters `json:"senders,omitempty"`
	Transactions   TokenCounters `json:"transactions,omitempty"`
}

type TokenCounters []TokenCounter

type TokenCounter struct {
	Token   Token `json:"token,omitempty"`
	Fails   int64 `json:"fails,string,omitempty"`
	Success int64 `json:"success,string,omitempty"`
	Total   int64 `json:"total,string,omitempty"`
}

type TransactionData struct {
	Results []*Counter `json:"results"`
}

type TransferData struct {
	Transfers    []CounterString `json:"transfers"`
	Transactions []CounterString `json:"transactions"`
}

type BalanceData struct {
	Supply  float64 `json:"supply,string"`
	Holders int64   `json:"holders,string"`
}

type BalanceDataResponse struct {
	Holders []BalanceData `json:"holders"`
}

type CounterString struct {
	Total   int64 `json:"total,string,omitempty"`
	Success int64 `json:"success,string,omitempty"`
	Fails   int64 `json:"fails,string,omitempty"`
}

type Counter struct {
	Total   int64 `json:"total,omitempty"`
	Success int64 `json:"success,omitempty"`
	Fails   int64 `json:"fails,omitempty"`
}

type HoldersResponse struct {
	HighHolders  []Holder     `json:"high_holders"`
	LowHolders   HoldersStats `json:"low_holders"`
	Holders      HoldersStats `json:"holders"`
	BurntBalance []Holder     `json:"burnt_balance"`
}

type HoldersStats []struct {
	Average           float64 `json:"average"`
	Count             int64   `json:"holders,string"`
	Median            float64 `json:"median"`
	StandardDeviation float64 `json:"standard_deviation"`
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
