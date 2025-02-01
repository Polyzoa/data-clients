package bitquery

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/elliotchance/pie/v2"
	"github.com/massigerardi/graphql"
	log "github.com/sirupsen/logrus"
)

type GraphqlClient interface {
	Run(background context.Context, req *graphql.Request, q interface{}) error
}

type Client struct {
	apiKey        string
	authorization string
	clientV1      graphql.Runner
	clientV2      graphql.Runner
}

func NewClientWithLog(apiKey string, auth string, log func(s string)) *Client {
	options := func(client *graphql.Client) {
		client.Log = log
	}
	clientV1 := graphql.NewClient("https://graphql.bitquery.io", options)
	clientV2 := graphql.NewClient("https://streaming.bitquery.io/graphql", options)
	return NewClient(apiKey, auth, clientV1, clientV2)
}

func NewClientDefault(apiKey string, auth string) *Client {
	clientV1 := graphql.NewClient("https://graphql.bitquery.io")
	clientV2 := graphql.NewClient("https://streaming.bitquery.io/graphql")
	return NewClient(apiKey, auth, clientV1, clientV2)
}

func NewClient(
	apiKey string,
	auth string,
	clientV1 GraphqlClient,
	clientV2 GraphqlClient,
) *Client {
	return &Client{
		apiKey:        apiKey,
		authorization: auth,
		clientV1:      clientV1,
		clientV2:      clientV2,
	}
}

func (c Client) runQuery(
	client GraphqlClient,
	query string,
	vars map[string]any,
	response interface{},
) error {
	req := graphql.NewRequest(query)
	for key, value := range vars {
		req.Var(key, value)
	}
	// req.Header.Set("X-API-KEY", c.apiKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	results := client.Run(context.Background(), req, &response)
	return results
}

func (c Client) GetStatsData(_ context.Context, chainName string, addresses []string) (*StatsData, error) {
	chain, err := GetChainV2(chainName)
	if err != nil {
		return nil, err
	}
	query := StatsQuery
	vars := map[string]any{
		"addresses": addresses,
		"network":   chain.Network,
	}
	var response Response[StatsData]
	err = c.runQuery(c.clientV2, query, vars, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c Client) GetStatsDataBefore(
	_ context.Context,
	chainName string,
	addresses []string,
	after time.Time,
	before time.Time,
) (*StatsData, error) {
	chain, err := GetChainV2(chainName)
	if err != nil {
		return nil, err
	}
	log.Debugf("Querying in interval %v - %v", after, before)
	query := StatsQueryInTime
	vars := map[string]any{
		"addresses": addresses,
		"network":   chain.Network,
		"after":     after.Format("2006-01-02T15:04:05Z"),
		"before":    before.Format("2006-01-02T15:04:05Z"),
	}
	var response Response[StatsData]
	err = c.runQuery(c.clientV2, query, vars, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c Client) GetContractData(_ context.Context, chainName string, addresses []string) (*ContractData, error) {
	chain, err := GetChainV1(chainName)
	if err != nil {
		return nil, err
	}
	query := ContractQuery
	vars := map[string]any{
		"addresses": addresses,
		"chain":     chain.Network,
	}
	var response Response[ContractData]
	err = c.runQuery(c.clientV1, query, vars, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c Client) GetHoldersData(
	_ context.Context,
	chainName string,
	address string,
	date time.Time,
	thresholdAmount ...float64,
) (
	*HoldersData, error,
) {
	chain, err := GetChainV2(chainName)
	if err != nil {
		return nil, err
	}
	query := HoldersQuery
	amount := 1000.0
	if len(thresholdAmount) > 0 {
		amount = pie.First(thresholdAmount)
	}
	amountAsString := strconv.FormatFloat(amount, 'f', -1, 64)
	vars := map[string]any{
		"address": address,
		"network": chain.Network,
		"date":    date.Format(time.DateOnly),
		"amount":  amountAsString,
	}
	var response Response[HoldersData]
	err = c.runQuery(c.clientV2, query, vars, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c Client) GetHistorySummary(chainName string, address string) (*Counter, error) {
	chain, err := GetChainV1(chainName)
	if err != nil {
		return nil, err
	}
	query := SuccessTransactionsQuery
	vars := map[string]any{
		"address": address,
		"chain":   chain,
	}
	var response Response[TransactionData]
	err = c.runQuery(c.clientV1, query, vars, &response)
	if err != nil {
		return nil, err
	}
	data := pie.First(response.Data.Results)
	return data, err
}

func (c Client) GetTransfersSummary(chainName string, address string) (*TransferData, error) {
	chain, err := GetChainV1(chainName)
	if err != nil {
		return nil, err
	}
	query := QueryTransactionStats
	vars := map[string]any{
		"address": address,
		"chain":   chain,
	}
	var response Response[TransferData]
	err = c.runQuery(c.clientV2, query, vars, &response)
	if err != nil {
		return nil, err
	}
	data := response.Data
	return &data, err
}

func (c Client) GetHolderBalance(chainName string, address string) (*BalanceData, error) {
	chain, err := GetChainV1(chainName)
	if err != nil {
		return nil, err
	}
	query := BalanceQuery
	vars := map[string]any{
		"address": address,
		"chain":   chain,
	}
	var response Response[BalanceDataResponse]
	err = c.runQuery(c.clientV2, query, vars, &response)
	if err != nil {
		return nil, err
	}
	data := response.Data
	if len(data.Holders) == 0 {
		return nil, errors.New("no data")
	}
	return &data.Holders[0], err
}
