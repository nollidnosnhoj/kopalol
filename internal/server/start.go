package server

import (
	"context"
	"net/http"
)

func (s *Server) Start(context context.Context) {
	server := http.Server{
		Addr:    ":8080",
		Handler: s.echo,
	}

	if err := server.ListenAndServe(); err != nil {
		s.echo.Logger.Fatal(err)
	}

	<-context.Done()
	err := server.Close()
	if err != nil {
		s.echo.Logger.Fatal(err)
	}
}
