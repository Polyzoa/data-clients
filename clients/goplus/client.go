package goplus

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"data-clients/clients/errors"
	"data-clients/clients/internal/httpclient"
	goplusModels "github.com/GoPlusSecurity/goplus-sdk-go/pkg/gen/models"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/massigerardi/go-commons/commons"
	log "github.com/sirupsen/logrus"
)

type TokenSecurity goplusModels.ResponseWrapperTokenSecurityResultAnon

type Client struct {
	client  httpclient.RetryableHttpClient
	maxWait int64
	minWait int64
}

func NewClient(client httpclient.RetryableHttpClient) *Client {
	maxWait := commons.GetInt("MAX_RETRY_WAIT", 120)
	minWait := commons.GetInt("MIN_RETRY_WAIT", 60)
	return &Client{
		client:  client,
		maxWait: maxWait,
		minWait: minWait,
	}
}

func (c Client) get(baseUrl string, response interface{}) error {
	req, err := retryablehttp.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return err
	}
	res, err := c.client.Do(req)
	if err != nil {
		log.Warnf("GET %v: %v", baseUrl, err)
		return err
	}
	if res.StatusCode != 200 {
		log.Warnf("GET %v: %v %v", baseUrl, res.StatusCode, res.Status)
		return fmt.Errorf("GET %v: %v", baseUrl, res.Status)
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Warnf("parsing body: %v: %v", res.Body, err)
		return err
	}
	return nil
}

func (c Client) GetTokenInfo(context context.Context, chainName string, address string) (*SecurityInfo, error) {
	chainId, err := getChainId(chainName)
	if err != nil {
		return nil, err
	}
	baseUrl := fmt.Sprintf("https://api.gopluslabs.io/api/v1/token_security/%v?contract_addresses=%v", chainId, address)
	var data Payload[SecurityInfoResult]
	count := 3
	var i int
	for i = 0; i < count; i++ {
		err = c.get(baseUrl, &data)
		if err != nil {
			return nil, fmt.Errorf("%v", err)
		}
		if data.Code == 1 {
			value, ok := data.Result[strings.ToLower(address)]
			if ok {
				return &value, nil
			}
		}
		if data.Code == 4029 {
			rest := math.Min(float64(c.maxWait), float64(c.minWait*int64(i)))
			time.Sleep(time.Duration(rest) * time.Second)
			log.Debugf("[%v] [%v] [attempt %v] retrying...", context.Value("id"), address, i+1)
		}
	}
	return nil, fmt.Errorf("no data fter %v attempts", i+1)
}

func getChainId(chainName string) (string, error) {
	if chainName == "ethereum" {
		return "1", nil
	}
	if chainName == "bsc" {
		return "56", nil
	}
	if chainName == "base" {
		return "8453", nil
	}
	return "", errors.ChainNotFoundError
}
