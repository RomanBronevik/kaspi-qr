package pg

import (
	"regexp"
	"time"
)

type transactionCtxKeyType bool

const (
	ErrPrefix         = "pg-error"
	transactionCtxKey = transactionCtxKeyType(true)
)

var defaultOptions = OptionsSt{
	Timezone:          "Asia/Almaty",
	MaxConns:          100,
	MinConns:          5,
	MaxConnLifetime:   30 * time.Minute,
	MaxConnIdleTime:   10 * time.Minute,
	HealthCheckPeriod: 20 * time.Second,
}

var (
	queryParamRegexp = regexp.MustCompile(`(?si)\$\{[^}]+\}`)
)
