package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type MessageRepository struct {
	c *redis.PubSubConn
}

func NewMsg(conn redis.Conn) *MessageRepository {
	return &MessageRepository{
		&redis.PubSubConn{Conn: conn},
	}
}

func (r *MessageRepository) Sub(channel string) error {
	return r.c.Subscribe(channel)
}

func (r *MessageRepository) Receive() {
	for r.c.Conn.Err() == nil {
		switch v := r.c.Receive().(type) {
		case redis.Message:
			fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
		case redis.Subscription:
			fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
		case error:
			fmt.Printf("err: %v", v)
		}
	}
}
