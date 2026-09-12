package transport

func WithAccept(value string) func(c *Client) {
	return func (c *Client){
		c.SetHeader("Accept", value) 
	}
}

func WithContentType(value string) func(c *Client) {
	return func (c *Client) {
		c.SetHeader("Content-Type", value)
	}
}
