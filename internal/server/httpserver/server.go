package httpserver

import (
	"context"

	"net/http"
	"redisjrpc/internal/config"
	"redisjrpc/pkg/logger"

	"github.com/getsentry/sentry-go"
)

type Server struct {
	Logger     logger.Logger
	httpServer *http.Server
}

func NewServer(cnfg config.HttpServer, router http.Handler, logger logger.Logger) *Server {
	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://f899ac51c7408446862547776c964df1@o4506480622567424.ingest.sentry.io/4506480648716288",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		logger.Errorf("Sentry initialization failed: %v", err)
	}

	srv := &http.Server{
		Addr:           cnfg.Host + ":" + cnfg.Port,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, //1 Mb
		ReadTimeout:    cnfg.ReadTimeout,
		WriteTimeout:   cnfg.WriteTimeout,
	}

	return &Server{
		Logger:     logger,
		httpServer: srv,
	}
}

func (s *Server) Start() error {
	err := s.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	s.Logger.Infof("Server starts at %s", s.httpServer.Addr)

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) GetAdress() string {
	return s.httpServer.Addr
}
