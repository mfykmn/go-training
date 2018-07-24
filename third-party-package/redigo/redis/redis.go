package redis

import (
	"github.com/gomodule/redigo/redis"
)

type Client struct {
	conn redis.Conn
}

// localhost:
func New(addr string) (*Client, error) {
	conn, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{
		conn: conn,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
