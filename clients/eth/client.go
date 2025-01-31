package eth

import (
	errorsUtils "errors"
	"time"

	"github.com/Polyzoa/data-clients/clients/errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

var urls = map[string]string{
	"base":     "https://rpc.ankr.com/base",
	"ethereum": "https://rpc.ankr.com/eth",
	"bsc":      "https://rpc.ankr.com/bsc",
	"matic":    "https://rpc.ankr.com/polygon",
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c Client) GetTokens(chain string, addresses []string) (map[string]bool, error) {
	url, ok := urls[chain]
	if !ok {
		return nil, errors.ChainNotFoundError
	}
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	var result = make(map[string]bool)
	for _, address := range addresses {
		var isToken bool = true
		for i := 0; i < 3; i++ {
			isToken, err = c.isToken(address, client)
			if err == nil {
				break
			}
			if errorsUtils.Is(err, errors.TooManyRequestsError) {
				log.Tracef("Too Many Requests, waiting 65 seconds [%v attempt]", i+1)
				time.Sleep(65 * time.Second)
				log.Tracef("Retrying...")
			} else {
				break
			}
		}
		result[address] = isToken
	}
	return result, nil
}

func (c Client) isToken(address string, client *ethclient.Client) (bool, error) {
	isToken, err := c.IsToken(address, client)
	if err != nil {
		if errorsUtils.As(err, &rpc.HTTPError{}) {
			var httpError rpc.HTTPError
			errorsUtils.As(err, &httpError)
			if httpError.StatusCode == 429 {
				return false, errors.TooManyRequestsError
			}
		} else {
			log.Warnf("error for address %s: %s", address, err)
			return false, err
		}
	}
	return isToken, nil
}

func (c Client) IsToken(address string, client *ethclient.Client) (bool, error) {
	tokenAddress := common.HexToAddress(address)
	token, err := NewToken(tokenAddress, client)
	if err != nil {
		return false, err
	}
	_, err = token.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return false, err
	}
	return true, nil
}
