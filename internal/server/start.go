package server

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nollidnosnhoj/kopalol/internal/log"
)

type Server struct {
	Addr    string
	Handler *echo.Echo
	logger  *slog.Logger
}

func NewServer(router *echo.Echo, logger *slog.Logger) *Server {
	return &Server{
		Addr:    ":8080",
		Handler: router,
		logger:  logger,
	}
}

func (s *Server) Start(context context.Context) {
	server := http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.LogErrorAndPanic(s.logger, err)
	}

	<-context.Done()
	err := server.Close()
	if err != nil {
		log.LogErrorAndPanic(s.logger, err)
	}
}
