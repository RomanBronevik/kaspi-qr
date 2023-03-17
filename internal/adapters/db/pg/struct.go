package pg

import (
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
