package bitquery

import (
	"context"
	"sync/atomic"

	"github.com/massigerardi/go-commons/commons"
	"github.com/massigerardi/graphql"
)

type Mock interface {
	Calls() int32
}

type MockGraphqlClient struct {
	Response []string
	error    []error
	counter  *atomic.Int32
}

func NewMockClient(response []string, error []error) *MockGraphqlClient {
	var counter atomic.Int32
	return &MockGraphqlClient{Response: response, error: error, counter: &counter}
}

func (m MockGraphqlClient) Calls() int32 {
	return m.counter.Load()
}

func (m MockGraphqlClient) Run(_ context.Context, _ *graphql.Request, resp interface{}) error {
	counter := m.counter.Load()
	defer m.counter.Add(1)
	if len(m.error) > int(counter) && m.error[counter] != nil {
		return m.error[counter]
	}
	err := commons.LoadFromJson(m.Response[counter], resp)
	return err
}
