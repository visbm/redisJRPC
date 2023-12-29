package service

import (
	"redisjrpc/internal/repository"
	"redisjrpc/pkg/logger"

	uuid "github.com/satori/go.uuid"
)

type ArticleService struct {
	articleRepository repository.ArticleRepository
	logger            logger.Logger
}

func NewArticleService(articleRepository repository.ArticleRepository, logger logger.Logger) *ArticleService {
	return &ArticleService{
		articleRepository: articleRepository,
		logger:            logger,
	}

}

func (s *ArticleService) GetArticles() ([]ArticleResponse, error) {
	articles, err := s.articleRepository.GetArticles()
	if err != nil {
		s.logger.Errorf("error while getting articles: %v", err)
		return nil, err
	}

	resp := MapArticlesToResponse(articles)
	return resp, nil
}

func (s *ArticleService) GetArticle(id string) (*ArticleResponse, error) {
	article, err := s.articleRepository.GetArticle(id)
	if err != nil {
		s.logger.Errorf("error while getting article: %v", err)
		return nil, err
	}
	resp := MapArticleToResponse(article)

	return &resp, nil
}

func (s *ArticleService) CreateArticle(req *ArticleRequest) (*ArticleResponse, error) {
	article := MapReqToArticle(*req)
	article.ID = uuid.NewV4().String()

	article, err := s.articleRepository.CreateArticle(article)
	if err != nil {
		s.logger.Errorf("error while creating article: %v", err)
		return nil, err
	}

	resp := MapArticleToResponse(article)

	return &resp, nil
}
