package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// redis ConnectionPooling
func newPool(addr string) *redis.Pool {
	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)

			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

type Client struct {
	*redis.Pool
}

func New(addr string) *Client {
	return &Client{
		newPool(addr),
	}
}
