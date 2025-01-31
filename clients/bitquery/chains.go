package bitquery

import (
	"github.com/Polyzoa/data-clients/clients/errors"
	"github.com/massigerardi/go-commons/commons"
)

type Chain struct {
	name    string
	network string
}

var ChainV2 = []Chain{
	{name: "ethereum", network: "eth"},
	{name: "base", network: "base"},
	{name: "optimism", network: "optimism"},
	{name: "bsc", network: "bsc"},
}

var ChainV1 = []Chain{
	{name: "ethereum", network: "ethereum"},
	{name: "matic", network: "matic"},
	{name: "optimism", network: "optimism"},
}

func GetChainV2(chainName string) (*Chain, error) {
	return getChain(ChainV2, chainName)
}

func GetChainV1(chainName string) (*Chain, error) {
	return getChain(ChainV1, chainName)
}

func getChain(chains []Chain, chainName string) (*Chain, error) {
	chain := commons.FindFirstUsing(chains, func(chain Chain) bool {
		return chain.name == chainName
	})
	if chain == nil {
		return nil, errors.ChainNotFoundError
	}
	return chain, nil
}
