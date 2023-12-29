package grpcserver

import (
	"context"
	"fmt"
	"net"
	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"
	article_proto "redisjrpc/protos/gen/proto/article"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

	"google.golang.org/grpc"
)

type grpcServer struct {
	gRPCServer *grpc.Server
	config     config.Grpc
	logger     logger.Logger

	article_proto.UnimplementedArticleServer
	articleService ArticleService
}

type ArticleServer interface {
	article_proto.ArticleServer
}

func NewGrpcServer(config config.Grpc, articleService ArticleService, logger logger.Logger) *grpcServer {
	gRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
			logging.UnaryServerInterceptor(logger),
		),
		grpc.ConnectionTimeout(config.Timeout),
	)

	return &grpcServer{
		gRPCServer:     gRPCServer,
		config:         config,
		articleService: articleService,
		logger:         logger,
	}
}

func (g *grpcServer) registerGrpcArticle(grpcServer *grpc.Server, article ArticleServer) {
	article_proto.RegisterArticleServer(grpcServer, article)
}

func (g *grpcServer) Start() error {
	g.registerGrpcArticle(g.gRPCServer, g)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", g.config.Port))
	if err != nil {
		return err
	}
	defer listener.Close()

	err = g.gRPCServer.Serve(listener)
	if err != nil {
		return err
	}
	g.logger.Infof("grpc Server starts at %s", g.config.Port)
	return nil
}

func (g *grpcServer) Shutdown(ctx context.Context) error {
	g.gRPCServer.GracefulStop()
	return nil
}

func (g *grpcServer) GetAdress() string {
	return fmt.Sprintf(":%s", g.config.Port)
}