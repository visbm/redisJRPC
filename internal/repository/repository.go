package repository

import (
	"errors"
	"redisjrpc/internal/config"
	"redisjrpc/internal/repository/models"
	"redisjrpc/internal/repository/pgsqldb"
	"redisjrpc/internal/repository/redisdb"
	"redisjrpc/pkg/logger"
)

type ArticleRepository interface {
	GetArticle(id string) (models.Article, error)
	GetArticles() ([]models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)

	Close() error
}

func NewArticleRepository(cfg config.Config, logger logger.Logger) (ArticleRepository, error) {
		
	switch cfg.DBType {
	case "redis":
		return repository.NewRedisDatabase(cfg.Redis, logger)
	case "pgSQL":
		return pgsqldb.NewPgSqlDB(cfg.PgSQL, logger)

	default:
		return nil, errors.New("invalid db type")
	}
}
