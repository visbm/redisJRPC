package repository

import (
	"github.com/go-redis/redis/v7"
	"redisjrpc/internal/config"	
	"redisjrpc/pkg/logger"
)

type RedisDatabase struct {
	Client *redis.Client
	logger logger.Logger
}


func NewRedisDatabase(cfg config.Redis, logger logger.Logger) (*RedisDatabase, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisDatabase{
		Client: rdb,
		logger: logger,
	}, nil
}

func (r *RedisDatabase) Close() error {
	return r.Client.Close()
}
