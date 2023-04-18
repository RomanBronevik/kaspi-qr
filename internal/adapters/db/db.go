package db

import (
	"context"
)

// Errors

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ErrNoRows = Err("no_rows")
)

// Interfaces

type DB interface {
	Connection
	Transaction
	HErr
}

type Connection interface {
	Exec(ctx context.Context, sql string, args ...any) error
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row

	ExecM(ctx context.Context, sql string, argMap map[string]interface{}) error
	QueryM(ctx context.Context, sql string, argMap map[string]interface{}) (Rows, error)
	QueryRowM(ctx context.Context, sql string, argMap map[string]interface{}) Row
}

type Transaction interface {
	TransactionFn(ctx context.Context, f func(context.Context) error) error
	RenewTransaction(ctx context.Context) (context.Context, error)
}

type HErr interface {
	HErr(err error) error
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
