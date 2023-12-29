package redisdb

import (
	"encoding/json"
	"redisjrpc/internal/repository/models"
	"redisjrpc/internal/database/redisdb"
	"redisjrpc/pkg/logger"
)


type redisArticleRepository struct {
	client *redisdb.RedisDatabase
	logger logger.Logger
} 


func NewRedisArticleRepository(cl *redisdb.RedisDatabase, logger logger.Logger) redisArticleRepository {
	return redisArticleRepository{
		client: cl,
		logger: logger,
	}
}


func (r redisArticleRepository) GetArticle(id string) (models.Article, error) {
	var a models.Article
	data, err := r.client.CL.Get(id).Bytes()
	if err != nil {
		r.logger.Errorf("%s", err)
		return a, err
	}
	err = json.Unmarshal(data, &a)
	if err != nil {
		r.logger.Errorf("%s", err)
		return a, err
	}

	return a, nil
}

func (r redisArticleRepository) GetArticles() ([]models.Article, error) {

	keys, err := r.client.CL.Keys("*").Result()
	if err != nil {
		r.logger.Errorf("%s", err)
		return nil, err
	}

	var articles []models.Article
	var a models.Article
	for _, key := range keys {
		data, err := r.client.CL.Get(key).Bytes()
		if err != nil {
			r.logger.Errorf("%s", err)
			return nil, err
		}

		err = json.Unmarshal(data, &a)
		if err != nil {
			r.logger.Errorf("%s", err)
			return nil, err
		}

		articles = append(articles, a)

	}
	return articles, nil
}

func (r redisArticleRepository) CreateArticle(a models.Article) (models.Article, error) {
	data, err := json.Marshal(a)
	if err != nil {
		r.logger.Errorf("%s", err)
		return a, err
	}

	err = r.client.CL.Set(a.ID, data, 0).Err()
	if err != nil {
		r.logger.Errorf("%s", err)
		return a, err
	}

	return a, nil
}
