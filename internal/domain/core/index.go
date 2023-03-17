package core

import (
	"sync"

	"kaspi-qr/internal/adapters/provider/kaspi"
	"kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/adapters/repo/pg"
)

type St struct {
	repo  repo.Repo
	kaspi *kaspi.St

	wg sync.WaitGroup
}

func New(repo *pg.St, kaspi *kaspi.St) *St {
	return &St{
		repo:  repo,
		kaspi: kaspi,
	}
}

func (s *St) WaitJobs() {
	s.wg.Wait()
}
