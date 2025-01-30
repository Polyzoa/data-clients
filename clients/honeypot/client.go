package honeypot

import (
	"context"
	"encoding/json"
	"fmt"

	"data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
)

type Client struct {
	client httpclient.RetryableHttpClient
}

func NewClient(client httpclient.RetryableHttpClient) *Client {
	return &Client{client: client}
}

func (c Client) GetSummary(context context.Context, address string) (*Summary, error) {
	baseUrl := fmt.Sprintf("https://api.honeypot.is/v2/IsHoneypot?address=%v", address)
	req, err := retryablehttp.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf(res.Status)
	}
	var response DataResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response.Summary, nil
}
