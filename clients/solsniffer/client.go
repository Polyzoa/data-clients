package solsniffer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	apiKey string
	client httpclient.RetryableHttpClient
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey, client: retryablehttp.NewClient()}
}

func NewClientWithRetryableHttpClient(apiKey string, client httpclient.RetryableHttpClient) *Client {
	return &Client{apiKey: apiKey, client: client}
}

func (c Client) get(path string, response interface{}) error {
	baseUrl := fmt.Sprintf("https://solsniffer.com/api/v2%v", path)
	req, err := retryablehttp.NewRequest(http.MethodGet, baseUrl, nil)
	if err != nil {
		return err
	}
	req.Header.Add("x-api-key", c.apiKey)
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

func (c Client) GetToken(address string) (*TokenData, error) {
	path := fmt.Sprintf("/token/%v", address)
	var response TokenData
	err := c.get(path, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
