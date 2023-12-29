package database

import (
	"errors"
	"redisjrpc/internal/config"
	"redisjrpc/internal/database/mongodb"
	"redisjrpc/internal/database/pgsqldb"
	"redisjrpc/internal/database/redisdb"
	"redisjrpc/pkg/logger"
)

type Database interface {
	Open(config config.Config, logger logger.Logger) error
	Close() error
}

func NewDatabase(cfg config.Config, logger logger.Logger) (Database, error) {
	var db Database
	switch cfg.DBType {
	case "redis":
		db = redisdb.NewRedisDatabase()
	case "pgSQL":
		db = pgsqldb.NewPgSqlDB()
	case "mongo":
		db = mongodb.NewMongoDB()
	default:
		return nil, errors.New("invalid db type")
	}
	err := db.Open(cfg, logger)
	if err != nil {
		return nil, err
	}
	return db, nil
}
