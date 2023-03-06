package pg

import (
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--

			continue

		}

		return nil
	}

	return
}

func (r *St) ErorrHandler(err error) error {
	var pgErr *pgconn.PgError
	if errors.Is(err, pgErr) {
		pgErr = err.(*pgconn.PgError)
		newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s", pgErr.Message, pgErr.Detail, pgErr.Where))
		fmt.Println(newErr)
		return newErr
	}
	return err
}
