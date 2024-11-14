package redis_api

import (
	"strconv"
	"web/internal/config"

	"github.com/go-redis/redis/v8"
)

const servicePrefix = "awesome_service." // наш префикс сервиса

type Client struct {
	cfg    config.Redis
	client *redis.Client
}

func New(cfg config.Redis) (*redis.Client, error) {
	client := &Client{}

	client.cfg = cfg

	redisClient := redis.NewClient(&redis.Options{
		Password:    cfg.Redis_password,
		Username:    cfg.Redis_user,
		Addr:        cfg.Redis_host + ":" + strconv.Itoa(cfg.Redis_port),
		DB:          0,
		DialTimeout: cfg.DialTimeout,
		ReadTimeout: cfg.ReadTimeout,
	})

	return redisClient, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
