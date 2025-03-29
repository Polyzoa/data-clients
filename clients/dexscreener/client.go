package dexscreener

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Polyzoa/data-clients/clients/errors"
	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/elliotchance/pie/v2"
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

func (c Client) GetLatestTokens(networks ...string) (Profiles, error) {
	baseUrl := "https://api.dexscreener.com/token-profiles/latest/v1"
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
	var response Profiles
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if len(networks) > 0 {
		response = pie.Filter(response, func(profile Profile) bool {
			return pie.Contains(networks, profile.ChainId)
		})
		if len(response) == 0 {
			return Profiles{}, nil
		}
	}
	return response, err
}

func (c Client) GetLatestBoostedTokens(networks ...string) (Profiles, error) {
	baseUrl := "https://api.dexscreener.com/token-boosts/latest/v1"
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
	var response Profiles
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if len(networks) > 0 {
		response = pie.Filter(response, func(profile Profile) bool {
			return pie.Contains(networks, profile.ChainId)
		})
		if len(response) == 0 {
			return Profiles{}, nil
		}
	}
	return response, err
}
func (c Client) GetTopBoostedTokens(networks ...string) (Profiles, error) {
	baseUrl := "https://api.dexscreener.com/token-boosts/top/v1"
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
	var response Profiles
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if len(networks) > 0 {
		response = pie.Filter(response, func(profile Profile) bool {
			return pie.Contains(networks, profile.ChainId)
		})
		if len(response) == 0 {
			return Profiles{}, nil
		}
	}
	return response, err
}
