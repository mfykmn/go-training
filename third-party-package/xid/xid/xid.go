package xid

import "github.com/rs/xid"

type Client struct {
	prefix string
}

func New(prefix string) *Client {
	return &Client{prefix}
}

func (s *Client) Create() string {
	guid := xid.New()
	return s.prefix + guid.String()
}
