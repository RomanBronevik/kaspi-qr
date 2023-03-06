package server

import (
	"kaspi-qr/internal/cns"
	"net/http"
)

import (
	"context"
)

type St struct {
	httpServer *http.Server
}

func (s *St) Run(port string, handler http.Handler) *St {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: cns.MaxHeaderBytes,
		ReadTimeout:    cns.ReadTimeout,
		WriteTimeout:   cns.WriteTimeout,
	}

	err := s.httpServer.ListenAndServe()

	if err != nil {
		return nil
	}

	return s
}

func (s *St) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
