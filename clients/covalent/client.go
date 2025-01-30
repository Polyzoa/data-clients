package covalent

import (
	"errors"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
	"github.com/covalenthq/covalent-api-sdk-go/services"
)

type Client struct {
	client CovalentClientType
}

func NewClient(apiKey string) *Client {
	client := covalentclient.CovalentClient(apiKey)
	covalentClient := CovalentClient{client: client}
	return &Client{client: &covalentClient}
}

func NewClientWithMock(mock CovalentClientType) *Client {
	return &Client{client: mock}
}

func (c Client) GetTransactionSummary(chainName string, address string) (*services.TransactionsSummary, error) {
	chain, err := getChain(chainName)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.GetTransactionSummary(chain, address)
	if err != nil {
		return nil, err
	}
	if resp.Error {
		return nil, errors.New(*resp.ErrorMessage)
	}
	if len(resp.Data.Items) == 0 {
		return nil, errors.New("no transactions found")
	}
	return &resp.Data.Items[0], nil
}

func getChain(chainName string) (chains.Chain, error) {
	switch chainName {
	case "ethereum":
		return chains.EthMainnet, nil
	case "eth":
		return chains.EthMainnet, nil
	case "bsc":
		return chains.BscMainnet, nil
	case "base":
		return chains.BaseMainnet, nil
	default:
		return "", errors.New("chain not found")
	}
}
