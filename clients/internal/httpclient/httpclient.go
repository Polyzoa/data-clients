package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"

	"github.com/hashicorp/go-retryablehttp"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	http.Client
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

type RetryableHttpClient interface {
	Do(request *retryablehttp.Request) (*http.Response, error)
}

// Mock Clients

type MockHttpClient struct {
	Responses   []string
	StatusCodes []int
	counter     *atomic.Int32
}

func NewMockHttpClient(responses []string, statusCodes []int) MockHttpClient {
	var counter atomic.Int32
	return MockHttpClient{
		Responses:   responses,
		StatusCodes: statusCodes,
		counter:     &counter,
	}
}

func (m MockHttpClient) Do(_ *http.Request) (*http.Response, error) {
	counter := m.counter.Load()
	response := &http.Response{
		Status:     fmt.Sprintf(http.StatusText(m.StatusCodes[counter])),
		StatusCode: m.StatusCodes[counter],
		Body:       io.NopCloser(strings.NewReader(m.Responses[counter])),
	}
	m.counter.Add(1)
	return response, nil
}

type Callable interface {
	Calls() int32
}

type MockRetryableHttpClient struct {
	Responses   []string
	StatusCodes []int
	counter     *atomic.Int32
	calls       *atomic.Int32
}

func (m MockRetryableHttpClient) Calls() int32 {
	return m.calls.Load()
}

func NewMockRetryableHttpClient(responses []string, statusCodes []int) MockRetryableHttpClient {
	var counter atomic.Int32
	var calls atomic.Int32
	return MockRetryableHttpClient{
		Responses:   responses,
		StatusCodes: statusCodes,
		counter:     &counter,
		calls:       &calls,
	}
}

func (m MockRetryableHttpClient) Do(_ *retryablehttp.Request) (*http.Response, error) {
	counter := m.counter.Load()
	var body io.ReadCloser
	status := m.StatusCodes[counter]
	if status == 200 {
		body = io.NopCloser(strings.NewReader(m.Responses[counter]))
	}
	response := &http.Response{
		Status:     fmt.Sprintf("Status %v", status),
		StatusCode: status,
		Body:       body,
	}
	m.counter.Add(1)
	m.calls.Add(1)
	return response, nil
}
