package app

import (
	"context"
	"redisjrpc/internal/config"
	"redisjrpc/internal/database"
	"redisjrpc/internal/server"

	"redisjrpc/internal/repository"
	"redisjrpc/internal/service"
	"redisjrpc/pkg/logger"
)

type App struct {
	logger     logger.Logger
	database   database.Database
	repository repository.ArticleRepository
	server     server.Server
}

func (a *App) StartApp(config config.Config, logger logger.Logger) error {
	a.logger = logger
	db, err := database.NewDatabase(config, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("db created ", config.DBType)
	a.database = db

	repo, err := repository.NewArticleRepository(config, db, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("repo created")
	a.repository = repo

	service := service.NewArticleService(repo, logger)
	logger.Info("service created")

	a.server = server.NewServer(config, service, logger)
	logger.Info("server created ", config.ServerType)

	go func() {
		if err := a.server.Start(); err != nil {
			logger.Fatal(err)
		}
	}()

	logger.Info("server started at ", a.server.GetAdress())
	return nil
}

func (a *App) Shutdown(ctx context.Context) error {
	err := a.server.Shutdown(context.Background())
	if err != nil {
		a.logger.Errorf("Error occurred on server shutting down: %s", err.Error())
	}

	err = a.database.Close()
	if err != nil {
		a.logger.Errorf("Error occurred on db connection close: %s", err.Error())
	}

	return nil
}
