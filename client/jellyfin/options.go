package jellyfin

import "context"

type Option func(*Client) error

func WithCredentials(ctx context.Context, username, password string) Option {
	return func(c *Client) error {
		return c.Auth.Login(ctx, username, password)
	}
}

func WithToken(token, userID string) Option {
	return func(c *Client) error {
		c.Auth.setToken(token)
		c.Auth.userID = userID
		return nil
	}
}
