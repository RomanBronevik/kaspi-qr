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
		MaxHeaderBytes: cns.MaxHeaderBytes, // 1 MB
		ReadTimeout:    cns.ReadTimeout,
		WriteTimeout:   cns.WriteTimeout,
	}

	s.httpServer.ListenAndServe()

	return s // Запускает бесконечный цикл и слушает все входящие запросы
}

func (s *St) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

//func NewHttpServer(port string, handler http.Handler) *http.Server {
//	httpServer := &http.Server{
//		Addr:           ":" + port,
//		Handler:        handler,
//		MaxHeaderBytes: cns.MaxHeaderBytes, // 1 MB
//		ReadTimeout:    cns.ReadTimeout,
//		WriteTimeout:   cns.WriteTimeout,
//	}
//	return httpServer
//}

//
