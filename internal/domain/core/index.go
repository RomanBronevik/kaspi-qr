package core

import (
	"kaspi-qr/internal/adapters/repo"
	"kaspi-qr/internal/adapters/repo/pg"
)

type St struct {
	repo repo.Repo
}

func New(repo *pg.St) *St {
	return &St{
		repo: repo,
	}
}
