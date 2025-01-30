package covalent

import (
	"encoding/json"

	"github.com/covalenthq/covalent-api-sdk-go/chains"
	"github.com/covalenthq/covalent-api-sdk-go/covalentclient"
	"github.com/covalenthq/covalent-api-sdk-go/services"
	"github.com/covalenthq/covalent-api-sdk-go/utils"
)

type CovalentClientType interface {
	GetTransactionSummary(chain chains.Chain, address string) (
		*utils.Response[services.TransactionsSummaryResponse], error,
	)
}

type CovalentClient struct {
	client *covalentclient.CovalentClientType
}

func (c CovalentClient) GetTransactionSummary(
	chain chains.Chain,
	address string,
) (*utils.Response[services.TransactionsSummaryResponse], error) {
	return c.client.TransactionService.GetTransactionSummary(chain, address)
}

type MockCovalentClient struct {
	response string
	err      error
}

func NewMockCovalentClient(response string, err error) *MockCovalentClient {
	return &MockCovalentClient{response: response, err: err}
}

func (m MockCovalentClient) GetTransactionSummary(
	_ chains.Chain,
	_ string,
) (*utils.Response[services.TransactionsSummaryResponse], error) {
	var summary services.TransactionsSummaryResponse
	err := json.Unmarshal([]byte(m.response), &summary)
	if err != nil {
		return nil, err
	}
	var message string
	if m.err != nil {
		message = m.err.Error()
	}
	return &utils.Response[services.TransactionsSummaryResponse]{
		Data:         &summary,
		Error:        m.err != nil,
		ErrorMessage: &message,
	}, nil
}
