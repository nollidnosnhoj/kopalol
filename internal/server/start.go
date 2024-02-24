package server

import (
	"context"
	"net/http"
)

func (s *Server) Start(context context.Context) {
	server := http.Server{
		Addr:    ":8080",
		Handler: s.Echo,
	}

	if err := server.ListenAndServe(); err != nil {
		s.Echo.Logger.Fatal(err)
	}

	<-context.Done()
	err := server.Close()
	if err != nil {
		s.Echo.Logger.Fatal(err)
	}
}
