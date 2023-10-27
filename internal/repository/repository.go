package repository

import (
	"errors"
	
	"redisjrpc/internal/database"
	"redisjrpc/internal/repository/models"
	"redisjrpc/internal/repository/mongodb"
	"redisjrpc/internal/repository/pgsqldb"
	"redisjrpc/internal/repository/redisdb" 
	mongodbDB "redisjrpc/internal/database/mongodb"
	pgsqldbDB "redisjrpc/internal/database/pgsqldb"
	redisdbDB "redisjrpc/internal/database/redisdb"
	"redisjrpc/pkg/logger"
	"redisjrpc/internal/config"
)

type ArticleRepository interface {
	GetArticle(id string) (models.Article, error)
	GetArticles() ([]models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
}

func NewArticleRepository(cfg config.Config,db database.Database, logger logger.Logger) (ArticleRepository, error) {
	switch cfg.DBType {
	 case "redis":
		return redisdb.NewRedisArticleRepository(db.(*redisdbDB.RedisDatabase), logger), nil
	case "pgSQL":
		return pgsqldb.NewArticlePgSqlRepository(db.(*pgsqldbDB.PgSqlDB), logger) , nil 
	case "mongo":
		return mongodb.NewArticleMongoRepository(db.(*mongodbDB.MongoDB), logger), nil
	default:
		return nil, errors.New("invalid db type")
	}
}
