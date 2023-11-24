package repository

import (
	"errors"

	"redisjrpc/internal/config"
	"redisjrpc/internal/database"
	"redisjrpc/internal/database/mongodb"
	"redisjrpc/internal/database/pgsqldb"
	"redisjrpc/internal/database/redisdb"
	"redisjrpc/internal/repository/models"
	mongodbRepo"redisjrpc/internal/repository/mongodb"
	pgsqldbRepo"redisjrpc/internal/repository/pgsqldb"
	redisdbRepo"redisjrpc/internal/repository/redisdb"
	"redisjrpc/pkg/logger"
)

type ArticleRepository interface {
	GetArticle(id string) (models.Article, error)
	GetArticles() ([]models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
}

func NewArticleRepository(cfg config.Config, db database.Database, logger logger.Logger) (ArticleRepository, error) {
	switch cfg.DBType {
	case "redis":
		return redisRepo.NewRedisArticleRepository(db.(*redisdb.RedisDatabase), logger), nil
	case "pgSQL":
		return pgsqlRepo.NewArticlePgSqlRepository(db.(*pgsqldb.PgSqlDB), logger), nil
	case "mongo":
		return mongoRepo.NewArticleMongoRepository(db.(*mongodb.MongoDB), logger), nil
	default:
		return nil, errors.New("invalid db type")
	}
}
