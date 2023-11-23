package handlers

import "redisjrpc/internal/service"

type ArticleRequest struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type ArticleResponse struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

type RespError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewRespError(err error, code int) RespError {
	return RespError{
		Message: err.Error(),
		Code:    code}

}

func mapArticleToResponse(a service.ArticleResponse) ArticleResponse {
	return ArticleResponse{
		ID:    a.ID,
		URL:   a.URL,
		Title: a.Title,
	}
}

func articleReq(a ArticleRequest) service.ArticleRequest {
	return service.ArticleRequest{
		URL:   a.URL,
		Title: a.Title,
	}
}

func mapArticlesResponse(articles []service.ArticleResponse) []ArticleResponse {
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
