// Package jellyfin abstracts the API for the jellyfin media platform.
package jellyfin

import (
	"net/url"

	"dolsh/client/jellyfin/transport"
)

const (
	clientName = "dolsh"
	deviceName = "dolsh"
	deviceID   = "dolsh"
	version    = "0.1.0"
)

type Client struct {
	Auth    *AuthService
	Library *LibraryService
	
	//internal
	transport *transport.Client
}

func New(baseURL *url.URL, opts ...Option) (*Client, error) {
	t := transport.New(
		baseURL, 
		transport.WithAccept("application/json"),
		transport.WithContentType("application/json"),
	)

	client := &Client{transport: t}
	client.Auth = &AuthService{Client: client}
	client.Library = &LibraryService{Client: client}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// Token returns the access token for the authenticated session, if any.
func (c *Client) Token() string {
	return c.Auth.Token()
}

// UserID returns the authenticated user's ID, if any.
func (c *Client) UserID() string {
	return c.Auth.UserID()
}
