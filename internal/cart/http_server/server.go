package http_server

import (
	"cart/internal/cart/api/cart_api"
	"cart/internal/cart/logger"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Run() error
	Stop() error
	AddHandlers(handlers []cart_api.HttpAPIHandler)
}

type server struct {
	ctx        context.Context
	mux        *http.ServeMux
	httpServer *http.Server
}

func NewServer(ctx context.Context, port string) Server {
	s := &server{
		ctx: ctx,
		mux: http.NewServeMux(),
	}

	s.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       10 * time.Second,
		Handler:           s.mux,
	}

	return s
}

func (s *server) Run() error {
	if err := s.httpServer.ListenAndServe(); err != nil {

		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("failed to start http server: %w", err)
		}
	}

	return nil
}

func (s *server) Stop() error {
	err := s.httpServer.Shutdown(s.ctx)
	if err != nil {
		logger.Errorf(s.ctx, "failed to stop http server: %v", err)
		return fmt.Errorf("failed to stop http server: %w", err)
	}
	logger.Info(s.ctx, "http server is stopped successfully")
	return nil
}

func (s *server) AddHandlers(handlers []cart_api.HttpAPIHandler) {
	for _, handler := range handlers {
		s.mux.HandleFunc(
			handler.Pattern,
			handler.Handler)
	}
}
