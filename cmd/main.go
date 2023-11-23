package main

import (
	"context"
	"os"
	"os/signal"

	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"
	"redisjrpc/cmd/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := logger.GetLogger()
	logger.Info("Starting application")

	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("config loaded db: %s , server: %s .", config.DBType, config.ServerType)

	application := app.App{}
	err = application.StartApp(*config, logger)
	if err != nil {
		logger.Fatal(err)
	}
	idleConnsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c

		err = application.Shutdown(context.Background())
		if err != nil {
			logger.Errorf("Error occurred on server shutting down: %s", err.Error())
		}

		logger.Info("shutting down")
		os.Exit(0)
	}()

	<-idleConnsClosed

}
