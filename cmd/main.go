package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"redisjrpc/internal/config"
	"redisjrpc/internal/handlers"
	"redisjrpc/internal/httpserver"
	"redisjrpc/internal/repository"
	"redisjrpc/internal/database"
	"redisjrpc/internal/service"
	"redisjrpc/pkg/logger"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("Starting application")

	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("config loaded ", config.DBType)

	db, err := database.NewDatabase(*config, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("db created ", db)


	repo, err := repository.NewArticleRepository(*config,db, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("repo created ")

	service := service.NewArticleService(repo, logger)
	articleHandler := handlers.NewArticleHandler(service, logger)
	handler := handlers.NewHandler(articleHandler)

	server := httpserver.NewServer(config.HttpServer, handler.InitRoutes(), logger)

	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		err = server.Shutdown(context.Background())
		if err != nil {
			logger.Errorf("Error occurred on server shutting down: %s", err.Error())
		}

		err = db.Close()
		if err != nil {
			logger.Errorf("Error occurred on db connection close: %s", err.Error())
		}

		logger.Info("shutting down")
		os.Exit(0)
	}()

	if err := server.Start(); err != http.ErrServerClosed {
		logger.Panicf("Error while starting server:%s", err)
	}
	<-idleConnsClosed

}
