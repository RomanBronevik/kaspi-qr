package pg

import (
	"github.com/jackc/pgx/v4"
	"kaspi-qr/internal/adapters/db"
	"time"
)

// Options

type OptionsSt struct {
	Dsn               string
	Timezone          string
	MaxConns          int32
	MinConns          int32
	MaxConnLifetime   time.Duration
	MaxConnIdleTime   time.Duration
	HealthCheckPeriod time.Duration
}

func (o *OptionsSt) mergeWithDefaults() {
	if o.Timezone == "" {
		o.Timezone = defaultOptions.Timezone
	}
	if o.MaxConns == 0 {
		o.MaxConns = defaultOptions.MaxConns
	}
	if o.MinConns == 0 {
		o.MinConns = defaultOptions.MinConns
	}
	if o.MaxConnLifetime == 0 {
		o.MaxConnLifetime = defaultOptions.MaxConnLifetime
	}
	if o.MaxConnIdleTime == 0 {
		o.MaxConnIdleTime = defaultOptions.MaxConnIdleTime
	}
	if o.HealthCheckPeriod == 0 {
		o.HealthCheckPeriod = defaultOptions.HealthCheckPeriod
	}
}

type rowsSt struct {
	pgx.Rows
	db db.HErr
}

func (o *rowsSt) Err() error {
	return o.db.HErr(o.Rows.Err())
}

func (o *rowsSt) Scan(dest ...any) error {
	return o.db.HErr(o.Rows.Scan(dest...))
}

type rowSt struct {
	pgx.Row
	db db.HErr
}

func (o *rowSt) Scan(dest ...any) error {
	return o.db.HErr(o.Row.Scan(dest...))
}
