package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Polyzoa/data-clients/clients/errors"
	"github.com/Polyzoa/data-clients/clients/internal/httpclient"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/massigerardi/go-commons/commons"
)

type Client struct {
	client   httpclient.RetryableHttpClient
	coins    Coins
	assets   Assets
	mutex    sync.Mutex
	initOnce sync.Once
}

func NewClient(httpclient httpclient.RetryableHttpClient) *Client {
	client := Client{client: httpclient, mutex: sync.Mutex{}, initOnce: sync.Once{}}
	client.Init()
	return &client
}

func (c *Client) Init() {
	c.initOnce.Do(func() {
		coins, err := c.getCoins()
		if err != nil {
			return
		}
		c.coins = coins
		assets, err := c.getAssets()
		if err != nil {
			return
		}
		c.assets = assets
	})

}

func (c *Client) getLocalFile(filepath string, result any) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &result)
}

func (c *Client) GetCoins() Coins {
	return c.coins
}

func (c *Client) GetAssets() Assets {
	return c.assets
}

func (c *Client) getCoins() (Coins, error) {
	filepath := commons.GetString("COINS_FILE", "coins.json")
	var coins Coins
	err := c.getLocalFile(filepath, &coins)
	if err == nil {
		return coins, nil
	}
	return c.getRemoteCoins()
}

func (c *Client) getAssets() (Assets, error) {
	filepath := commons.GetString("ASSETS_FILE", "assets.json")
	var assets Assets
	err := c.getLocalFile(filepath, &assets)
	if err == nil {
		return assets, nil
	}
	return c.getRemoteAssets()
}

func (c *Client) GetAsset(chainName string) (*Asset, error) {
	asset := commons.FindFirstUsing(c.assets, func(value Asset) bool {
		return strings.ToLower(value.Name) == strings.ToLower(chainName) ||
			strings.ToLower(value.Shortname) == strings.ToLower(chainName)
	})
	if asset != nil {
		return asset, nil
	}
	return nil, errors.ChainNotFoundError
}

func (c *Client) GetCoin(context context.Context, chainName string, address string) (*FullCoin, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	chain, err := c.GetAsset(chainName)
	if err != nil {
		return nil, err
	}
	coin := commons.FindFirstUsing(c.coins, func(coin Coin) bool {
		coinAddress, ok := coin.Platforms[chain.Id]
		return ok && coinAddress == address
	})
	if coin == nil {
		return nil, errors.CoinNotFoundError
	}
	return c.getCoinById(context, coin.Id)
}

func (c *Client) getCoinById(_ context.Context, coinId string) (*FullCoin, error) {
	baseUrl := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%v", coinId)
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
	var response FullCoin
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) getRemoteCoins() (Coins, error) {
	filepath := commons.GetString("COINS_FILE", "coins.json")
	baseUrl := "https://api.coingecko.com/api/v3/coins/list?include_platform=true"
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
	var response Coins
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	commons.ToJsonFile(filepath, response)
	return response, nil
}

func (c *Client) getRemoteAssets() (Assets, error) {
	filepath := commons.GetString("ASSETS_FILE", "assets.json")
	baseUrl := "https://api.coingecko.com/api/v3/asset_platforms"
	api_key := commons.GetString("COINGECKO_API_KEY", "")
	req, err := retryablehttp.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("x-cg-demo-api-key", api_key)
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf(res.Status)
	}
	var response Assets
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	commons.ToJsonFile(filepath, response)
	return response, nil
}
