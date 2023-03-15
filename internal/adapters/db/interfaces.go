package db

import (
	"context"
)

type DB interface {
	Connection
	Transaction

	HErr(err error) error
}

type Connection interface {
	Exec(ctx context.Context, sql string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

type Transaction interface {
	TransactionFn(ctx context.Context, f func(context.Context) error) error
	RenewTransaction(ctx context.Context) (context.Context, error)
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Row interface {
	Scan(dest ...any) error
}
