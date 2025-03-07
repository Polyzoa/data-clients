package goplus

import (
	"encoding/json"
	"errors"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
)

type SolanaClient struct {
	client httpclient.RetryableHttpClient
}

func NewSolanaClient() *SolanaClient {
	return &SolanaClient{client: retryablehttp.NewClient()}
}

func (c SolanaClient) GetSecurity(address string) (*SolanaSecurityInfo, error) {
	baseUrl := "https://api.gopluslabs.io/api/v1/solana/token_security"
	req, err := retryablehttp.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("contract_addresses", address)
	req.URL.RawQuery = query.Encode()
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	var response Payload[SolanaSecurityInfoResult]
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	result := response.Result[address]
	return &result, nil
}
