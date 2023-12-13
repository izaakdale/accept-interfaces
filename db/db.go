package db

import "github.com/go-redis/redis"

type RedisClient interface {
	Ping() *redis.StatusCmd
}

type connection struct {
	client *redis.Client
}

func New(cli *redis.Client) (*connection, error) {
	return &connection{cli}, nil
}

func (c *connection) Ping() error {
	return c.client.Ping().Err()
}
