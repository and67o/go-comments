package redis

import (
	"github.com/and67o/go-comments/internal/configuration"
	"github.com/and67o/go-comments/internal/interfaces"
	"github.com/go-redis/redis"
	"net"
	"time"
)

type Redis struct {
	*redis.Client
}

func New(conf configuration.Redis) (interfaces.Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(conf.Host, conf.Port),
		Password: conf.Password,
	})

	_, err := client.Ping().Result()

	if err != nil {
		return nil, err
	}

	return &Redis{client}, nil
}

func (r *Redis) Get(key string) (interface{},error) {
	return nil, nil
}

func (r *Redis) Set(key string, value interface{}, expiration time.Duration) error {
	return nil
}
