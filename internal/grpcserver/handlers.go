package grpcserver

import (
	"context"
	"redisjrpc/internal/service"
	article_proto "redisjrpc/protos/gen/proto/article"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArticleService interface {
	GetArticle(id string) (*service.ArticleResponse, error)
	GetArticles() ([]service.ArticleResponse, error)
	CreateArticle(article *service.ArticleRequest) (*service.ArticleResponse, error)
}

func (s *grpcServer) GetArticle(ctx context.Context, id *article_proto.GetArticleID) (*article_proto.ArticleResponse, error) {

	article, err := s.articleService.GetArticle(id.Id)
	if err != nil {
		return &article_proto.ArticleResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	return &article_proto.ArticleResponse{
		Id:    article.ID,
		Title: article.Title,
		Url:   article.URL,
	}, nil

}

func (s *grpcServer) CreateArticle(ctx context.Context, article *article_proto.ArticleRequest) (*article_proto.ArticleResponse, error) {

	resp, err := s.articleService.CreateArticle(&service.ArticleRequest{
		Title: article.Title,
		URL:   article.Url,
	})
	if err != nil {
		return &article_proto.ArticleResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &article_proto.ArticleResponse{
		Id:    resp.ID,
		Title: resp.Title,
		Url:   resp.URL,
	}, nil

}

func (s *grpcServer) GetArticles(ctx context.Context, id *article_proto.Empty) (*article_proto.ArticleArray, error) {
	articles, err := s.articleService.GetArticles()
	if err != nil {
		return &article_proto.ArticleArray{}, status.Error(codes.InvalidArgument, err.Error())
	}

	var articleArray []*article_proto.ArticleResponse
	for _, article := range articles {
		articleArray = append(articleArray, &article_proto.ArticleResponse{
			Id:    article.ID,
			Title: article.Title,
			Url:   article.URL,
		})
	}

	return &article_proto.ArticleArray{
		Values: articleArray,
	}, nil

}
