package bitquery

import (
	"github.com/Polyzoa/data-clients/clients/errors"
	"github.com/massigerardi/go-commons/commons"
)

type Chain struct {
	Name    string
	Network string
}

var ChainV2 = []Chain{
	{Name: "ethereum", Network: "eth"},
	{Name: "base", Network: "base"},
	{Name: "optimism", Network: "optimism"},
	{Name: "bsc", Network: "bsc"},
}

var ChainV1 = []Chain{
	{Name: "ethereum", Network: "ethereum"},
	{Name: "matic", Network: "matic"},
	{Name: "optimism", Network: "optimism"},
}

func GetChainV2(chainName string) (*Chain, error) {
	return getChain(ChainV2, chainName)
}

func GetChainV1(chainName string) (*Chain, error) {
	return getChain(ChainV1, chainName)
}

func getChain(chains []Chain, chainName string) (*Chain, error) {
	chain := commons.FindFirstUsing(chains, func(chain Chain) bool {
		return chain.Name == chainName
	})
	if chain == nil {
		return nil, errors.ChainNotFoundError
	}
	return chain, nil
}
