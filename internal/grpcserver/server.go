package grpcserver

import (
	"fmt"
	"net"
	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"
	article_proto "redisjrpc/protos/gen/proto/article"

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

func NewGrpcServer(config config.Grpc, articleService ArticleService) *grpcServer {
	gRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(),
		),
		grpc.ConnectionTimeout(config.Timeout),
	)

	return &grpcServer{
		gRPCServer:     gRPCServer,
		config:         config,
		articleService: articleService,
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
	err = g.gRPCServer.Serve(listener)
	if err != nil {
		return err
	}

	g.logger.Infof("grpc Server starts at %s", g.config.Port)

	return nil
}

func (g *grpcServer) Stop() {
	g.gRPCServer.Stop()
}
