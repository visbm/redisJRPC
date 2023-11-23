package handlers

import (
	"fmt"
	"net/http"
	"redisjrpc/internal/service"
	"redisjrpc/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type ArticleService interface {
	GetArticle(id string) (*service.ArticleResponse, error)
	GetArticles() ([]service.ArticleResponse, error)
	CreateArticle(article *service.ArticleRequest) (*service.ArticleResponse, error)
}

type handler struct {
	service ArticleService
	logger  logger.Logger
}

func NewArticleHandler(service ArticleService, logger logger.Logger) ArticleHandler {
	return &handler{
		service: service,
		logger:  logger,
	}
}

func (h handler) SaveArticle(w http.ResponseWriter, r *http.Request) {
	var req ArticleRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		h.logger.Errorf("%s", err)
		render.JSON(w, r, NewRespError(err, http.StatusBadRequest))
		return
	}

	createArticle := articleReq(req)
	resp, err := h.service.CreateArticle(&createArticle)
	if err != nil {
		h.logger.Errorf("%s", err)
		render.JSON(w, r, NewRespError(err, http.StatusBadRequest))
		return
	}

	render.JSON(w, r, resp)
}

func (h handler) GetArticle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.logger.Info("alias is empty")
		render.JSON(w, r, NewRespError(fmt.Errorf("alias is empty"), http.StatusBadRequest))
		return
	}

	article, err := h.service.GetArticle(id)
	if err != nil {
		h.logger.Errorf("%s", err)
		render.JSON(w, r, NewRespError(err, http.StatusBadRequest))
		return
	}

	resp := mapArticleToResponse(*article)
	render.JSON(w, r, resp)
}

func (h handler) GetArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetArticles()
	if err != nil {
		h.logger.Errorf("%s", err)
		render.JSON(w, r, NewRespError(err, http.StatusBadRequest))

	}
	resp := mapArticlesResponse(articles)
	render.JSON(w, r, resp)
}
