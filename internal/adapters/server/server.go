package server

import (
	"context"
	"net/http"
	"time"

	"kaspi-qr/internal/adapters/logger"
)

const (
	ReadHeaderTimeout = 10 * time.Second
	ReadTimeout       = 2 * time.Minute
	MaxHeaderBytes    = 300 * 1024
)

type St struct {
	lg   logger.Lite
	addr string

	server *http.Server
	eChan  chan error
}

func Start(lg logger.Lite, addr string, handler http.Handler) *St {
	s := &St{
		lg:   lg,
		addr: addr,
		server: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadHeaderTimeout: ReadHeaderTimeout,
			ReadTimeout:       ReadTimeout,
			MaxHeaderBytes:    MaxHeaderBytes,
		},
		eChan: make(chan error, 1),
	}

	s.lg.Infow("Start rest-api", "addr", s.server.Addr)

	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.lg.Errorw("Http server closed", err)
			s.eChan <- err
		}
	}()

	return s
}

func (s *St) Wait() <-chan error {
	return s.eChan
}

func (s *St) Shutdown(timeout time.Duration) bool {
	defer close(s.eChan)

	ctx, ctxCancel := context.WithTimeout(context.Background(), timeout)
	defer ctxCancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		s.lg.Errorw("Fail to shutdown http-api", err, "addr", s.addr)
		return false
	}

	return true
}
