package redisdb

import (
	"github.com/go-redis/redis/v7"
	"redisjrpc/internal/config"	
	"redisjrpc/pkg/logger"
)

type RedisDatabase struct {
	CL *redis.Client
	logger logger.Logger
}

func NewRedisDatabase() *RedisDatabase{
	return &RedisDatabase{}
}

func (m *RedisDatabase) Open(cfg config.Config, logger logger.Logger) error {
    opts, err := redis.ParseURL(cfg.Redis.Url)
    if err != nil {
        return  err
    }
    rdb := redis.NewClient(opts)

	_, err = rdb.Ping().Result()
	if err != nil {
		return  err
	}

	m.CL = rdb
	m.logger = logger
	return nil
}

func (r *RedisDatabase) Close() error {
	return r.CL.Close()
}
