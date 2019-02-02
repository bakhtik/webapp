package models

import "github.com/go-redis/redis"

type Cache interface {
	Increment(key string) (result int64, err error)
}

type Client struct {
	*redis.Client
}

func NewClient(addr string) (client *Client) {
	return &Client{
		redis.NewClient(&redis.Options{
			Addr: addr, // "redis:6379" when in container
		}),
	}
}

func (c *Client) Increment(key string) (result int64, err error) {
	return c.Client.Incr(key).Result()
}
