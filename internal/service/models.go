package service

import "redisjrpc/internal/repository/models"

type ArticleRequest struct {
	URL   string
	Title string
}

type ArticleResponse struct {
	ID    string
	URL   string
	Title string
}

func MapArticleToResponse(a models.Article) ArticleResponse {
	return ArticleResponse{
		ID:    a.ID,
		URL:   a.URL,
		Title: a.Title,
	}
}

func MapReqToArticle(a ArticleRequest) models.Article {
	return models.Article{
		URL:   a.URL,
		Title: a.Title,
	}
}

func MapArticlesToResponse(articles []models.Article) []ArticleResponse {
	var resp []ArticleResponse
	for _, a := range articles {
		r := ArticleResponse{
			ID:    a.ID,
			URL:   a.URL,
			Title: a.Title,
		}
		resp = append(resp, r)
	}
	return resp
}
