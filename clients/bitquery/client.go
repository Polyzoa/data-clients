package bitquery

import (
	"context"
	"errors"
	"fmt"

	"github.com/massigerardi/graphql"
)

type GraphqlClient interface {
	Run(background context.Context, req *graphql.Request, q interface{}) error
}

type Client struct {
	apiKey        string
	authorization string
}

func NewClient(
	apiKey string,
	auth string,
) *Client {
	return &Client{
		apiKey:        apiKey,
		authorization: auth,
	}
}

func (c Client) RunQuery(
	query Query,
	response interface{},
) error {
	if query.Params == nil || len(query.Params) == 0 {
		return errors.New("no parameters provided")
	}
	req := graphql.NewRequest(query.Query)
	for key, value := range query.Params {
		req.Var(key, value)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	client := getGraphqlClient(query.Url)
	var err error
	err = client.Run(context.Background(), req, response)
	return err
}
