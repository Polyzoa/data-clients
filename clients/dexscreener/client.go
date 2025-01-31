package dexscreener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Polyzoa/data-clients/clients/errors"
	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
)

type Client struct {
	client httpclient.HttpClient
}

func NewClient(client httpclient.HttpClient) *Client {
	return &Client{client: client}
}

func (c Client) GetPairs(addresses []string) (Pairs, error) {
	baseUrl := fmt.Sprintf("https://api.dexscreener.com/latest/dex/tokens/%v", strings.Join(addresses, ","))
	req, err := http.NewRequest("GET", baseUrl, nil)
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
	var response PairsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if len(response.Pairs) == 0 {
		return nil, errors.NotFoundError
	}
	return response.Pairs, nil

}
