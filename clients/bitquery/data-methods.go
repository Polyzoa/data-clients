package bitquery

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/elliotchance/pie/v2"
	log "github.com/sirupsen/logrus"
)

func (c Client) GetStatsData(_ context.Context, chainName string, addresses []string) (*StatsData, error) {
	chain, err := GetChainV2(chainName)
	if err != nil {
		return nil, err
	}
	query := StatsQuery.Clone()
	query.Params = map[string]any{
		"addresses": addresses,
		"network":   chain.Network,
	}
	var response Response[StatsData]
	err = c.RunQuery(query, &response)
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
	query := StatsQueryInTime.Clone()
	query.Params = map[string]any{
		"addresses": addresses,
		"network":   chain.Network,
		"after":     after.Format("2006-01-02T15:04:05Z"),
		"before":    before.Format("2006-01-02T15:04:05Z"),
	}
	var response Response[StatsData]
	err = c.RunQuery(query, &response)
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
	query := ContractQuery.Clone()
	query.Params = map[string]any{
		"addresses": addresses,
		"chain":     chain.Network,
	}
	var response Response[ContractData]
	err = c.RunQuery(query, &response)
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
	query := HoldersQuery.Clone()
	amount := 1000.0
	if len(thresholdAmount) > 0 {
		amount = pie.First(thresholdAmount)
	}
	amountAsString := strconv.FormatFloat(amount, 'f', -1, 64)
	query.Params = map[string]any{
		"address": address,
		"network": chain.Network,
		"date":    date.Format(time.DateOnly),
		"amount":  amountAsString,
	}
	var response Response[HoldersData]
	err = c.RunQuery(query, &response)
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
	query := SuccessTransactionsQuery.Clone()
	query.Params = map[string]any{
		"address": address,
		"chain":   chain,
	}
	var response Response[TransactionData]
	err = c.RunQuery(query, &response)
	if err != nil {
		return nil, err
	}
	data := pie.First(response.Data.Results)
	return data, nil
}

func (c Client) GetTransfersSummary(chainName string, address string) (*TransferData, error) {
	chain, err := GetChainV1(chainName)
	if err != nil {
		return nil, err
	}
	query := QueryTransactionStats.Clone()
	query.Params = map[string]any{
		"address": address,
		"chain":   chain,
	}
	var response Response[TransferData]
	err = c.RunQuery(query, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data, nil
}

func (c Client) GetHolderBalance(chainName string, address string) (*BalanceData, error) {
	chain, err := GetChainV1(chainName)
	if err != nil {
		return nil, err
	}
	query := BalanceQuery.Clone()
	query.Params = map[string]any{
		"address": address,
		"chain":   chain,
	}
	var response Response[BalanceDataResponse]
	err = c.RunQuery(query, &response)
	if err != nil {
		return nil, err
	}
	data := response.Data
	if len(data.Holders) == 0 {
		return nil, errors.New("no data")
	}
	return &data.Holders[0], err
}
