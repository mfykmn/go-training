package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type MessageRepository struct {
	c *redis.PubSubConn
}

func NewMsg(client *Client) *MessageRepository {
	return &MessageRepository{
		&redis.PubSubConn{Conn: client.conn},
	}
}

func (r *MessageRepository) Sub(channel string) error {
	return r.c.Subscribe(channel)
}

func (r *MessageRepository) Pub(channel, msg string) error {
	err := r.c.Conn.Send("PUBLISH", channel, msg)
	if err != nil {
		return err
	}
	return r.c.Conn.Flush()
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

func (r *MessageRepository) Close() error {
	return r.Close()
}
