package mysqldb

import (
	"redisjrpc/internal/repository/models"
)

func (r MysqlRepository) GetArticle(id string) (models.Article, error) {
	return models.Article{}, nil
}

func (r MysqlRepository) GetArticles() ([]models.Article, error) {
	return []models.Article{
		{
			ID:    "1",
			URL:   "https://www.example.com",
			Title: "Example",
		},
	}, nil
}

func (r MysqlRepository) CreateArticle(a models.Article) (models.Article, error) {
	return a, nil
}
