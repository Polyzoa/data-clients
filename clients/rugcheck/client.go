package rugcheck

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	client httpclient.RetryableHttpClient
}

func NewClient() *Client {
	return &Client{client: retryablehttp.NewClient()}
}

func NewClientWithRetryableHttpClient(client httpclient.RetryableHttpClient) *Client {
	return &Client{client: client}
}

func (c Client) get(path string, response interface{}) error {
	baseUrl := fmt.Sprintf("https://api.rugcheck.xyz/v1%v", path)
	req, err := retryablehttp.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf(res.Status)
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	return err
}

func (c Client) GetLatestTokens() ([]*Token, error) {
	path := "/stats/new_tokens"
	var response []*Token
	err := c.get(path, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c Client) GetTokenSummary(address string) (*TokenData, error) {
	path := fmt.Sprintf("/tokens/%v/report/summary", address)
	var response TokenData
	err := c.get(path, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
