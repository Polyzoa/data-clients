package coingecko

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"data-clients/clients/errors"
	"data-clients/clients/internal/httpclient"
	"github.com/elliotchance/pie/v2"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/massigerardi/go-commons/commons"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	client   httpclient.RetryableHttpClient
	coins    Coins
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
	})
}

func (c *Client) getLocalFile(filepath string) (Coins, error) {
	var coins Coins
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Warnf("getting local coins file failed: %v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &coins)
	if err != nil {
		log.Warnf("Unmarshal local coins file failed: %v", err)
		return nil, err
	}
	return coins, nil
}

func (c *Client) GetCoins() Coins {
	return c.coins
}

func (c *Client) getCoins() (Coins, error) {
	filepath := commons.GetString("COINS_FILE", "coins.json")
	coins, err := c.getLocalFile(filepath)
	if err == nil {
		return coins, nil
	}
	return c.getRemoteCoins()
}

func (c *Client) GetCoin(context context.Context, chainName string, address string) (*FullCoin, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	chain := getChain(chainName)
	if chain == nil {
		return nil, errors.NotFoundError
	}
	founds := pie.Filter(c.coins, func(coin Coin) bool {
		coinAddress, ok := coin.Platforms[chain.name]
		return ok && coinAddress == address
	})
	coin := pie.First(founds)
	if coin.Id == "" {
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
