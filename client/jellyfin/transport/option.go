package transport

import (
	"strconv"
	"net/url"
)


type clientOption func(*Client)

func WithAccept(value string) clientOption {
	return func (c *Client){
		c.SetHeader("Accept", value) 
	}
}

func WithContentType(value string) clientOption {
	return func (c *Client) {
		c.SetHeader("Content-Type", value)
	}
}

type queryOption func(params *url.Values)

func WithLimit(limit int) queryOption {
	return func(params *url.Values) {
		params.Add("Limit", strconv.Itoa(limit))
	}
}

func WithParentID(id string) queryOption {
	return func(params *url.Values) {
		params.Add("ParentId", id)
	}
}

func WithMusicAlbumType () queryOption {
	return func(params *url.Values) {
		params.Add("IncludeItemTypes", "MusicAlbum")
	}
}
