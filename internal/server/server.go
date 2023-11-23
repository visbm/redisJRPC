package server

import (
	"context"
	"redisjrpc/internal/config"
	"redisjrpc/internal/server/grpcserver"
	"redisjrpc/internal/server/httpserver"
	"redisjrpc/internal/server/httpserver/handlers"
	"redisjrpc/internal/service"
	"redisjrpc/pkg/logger"
)

type Server interface {
	Start() error
	Shutdown(context.Context) error
	GetAdress() string
}

func NewServer(cfg config.Config, service *service.ArticleService, logger logger.Logger) Server {
	switch cfg.ServerType {
	case "http":
		return getHttpServer(cfg, service, logger)
	case "grpc":
		return getGrpcServer(cfg, service, logger)
	default:
		logger.Fatal("unknown server type")
		return nil
	}

}

func getHttpServer(cfg config.Config, service *service.ArticleService, logger logger.Logger) Server {
	articleHandler := handlers.NewArticleHandler(service, logger)
	handler := handlers.NewHandler(articleHandler)
	logger.Info("handler created")

	return httpserver.NewServer(cfg.HttpServer, handler.InitRoutes(), logger)

}

func getGrpcServer(cfg config.Config, service *service.ArticleService, logger logger.Logger) Server {
	return grpcserver.NewGrpcServer(cfg.Grpc, service, logger)
}
