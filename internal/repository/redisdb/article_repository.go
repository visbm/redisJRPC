package repository

import (
	"encoding/json"
	"redisjrpc/internal/repository/models"
)

func (r RedisDatabase) GetArticle(id string) (models.Article, error) {
	var a models.Article
	data, err := r.Client.Get(id).Bytes()
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

func (r RedisDatabase) GetArticles() ([]models.Article, error) {

	keys, err := r.Client.Keys("*").Result()
	if err != nil {
		r.logger.Errorf("%s", err)
		return nil, err
	}

	var articles []models.Article
	var a models.Article
	for _, key := range keys {
		data, err := r.Client.Get(key).Bytes()
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

func (r RedisDatabase) CreateArticle(a models.Article) (models.Article, error) {
	data, err := json.Marshal(a)
	if err != nil {
		r.logger.Errorf("%s", err)
		return a, err
	}

	err = r.Client.Set(a.ID, data, 0).Err()
	if err != nil {
		r.logger.Errorf("%s", err)
		return a, err
	}

	return a, nil
}
