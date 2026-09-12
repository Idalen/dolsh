package transport

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient http.Client
	baseURL    url.URL
	headers    http.Header
}

type option func(*Client)

func New(baseURL *url.URL, opts ...option) *Client {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	client := &Client{
		httpClient: httpClient,
		baseURL:    *baseURL,
		headers:    http.Header{},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// SetHeader sets a header sent on every request.
func (c *Client) SetHeader(key, value string) {
	c.headers.Set(key, value)
}

func (c *Client) Request[Request any, Response any](
	ctx context.Context,
	method string,
	path string,
	body Request,
) (Response, error) {
	var res Response

	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return res, err
	}

	res, err = c.do[Response](req)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c *Client) do[T any](req *http.Request) (T, error) {
	var data T

	res, err := c.httpClient.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return data, newAPIError(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, err
	}

	return data, nil
}

func (c *Client) newRequest[T any](
	ctx context.Context,
	method string,
	path string,
	body T,
) (*http.Request, error) {
	var buf bytes.Buffer

	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	u := c.baseURL.JoinPath(path)

	req, err := http.NewRequestWithContext(
		ctx,
		method,
		u.String(),
		&buf,
	)
	if err != nil {
		return nil, err
	}

	for key, values := range c.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	return req, nil
}
