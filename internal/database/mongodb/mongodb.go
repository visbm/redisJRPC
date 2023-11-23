package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"
	"time"
)

type MongoDB struct {
	Cl *mongo.Client	
	logger logger.Logger
}

func NewMongoDB() *MongoDB {
	return &MongoDB{}
}

func (m *MongoDB) Open(cfg config.Config, logger logger.Logger)  error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Mongo.Username, cfg.Mongo.Password, cfg.Mongo.Host, cfg.Mongo.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logger.Panicf("mongodb open error:%s", err)
		return  err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Panicf("mongodb ping error:%s", err)
		return  err
	}

	m.Cl = client
	m.logger = logger
	return nil
}

func (m *MongoDB) Close() error {
	return m.Cl.Disconnect(context.Background())
}
