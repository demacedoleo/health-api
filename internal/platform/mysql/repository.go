package mysql

import "C"
import (
	"context"
	"database/sql"
)

const (
	noResult = 0
)

type Repository interface {
	Begin(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	Save(ctx context.Context, sql string, args ...interface{}) (int64, error)
	Update(ctx context.Context, sql string, args ...interface{}) (int64, error)
	Fetch(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error)
}

func NewRepository(db *sql.DB) *persistence {
	if db != nil {
		return &persistence{
			client: &client{db: db},
		}
	}

	return &persistence{
		client: NewDefaultClient(),
	}
}

type persistence struct {
	*client
}

func (p *persistence) Begin(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error) {
	return p.db.BeginTx(ctx, options)
}

func (p *persistence) Save(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	stmt, err := p.db.Prepare(sql)
	if err != nil {
		return noResult, err
	}

	defer func() {
		_ = stmt.Close()
	}()

	result, err := p.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return noResult, err
	}

	return result.LastInsertId()
}

func (p *persistence) Update(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	stmt, err := p.db.Prepare(sql)
	if err != nil {
		return noResult, err
	}

	defer func() {
		_ = stmt.Close()
	}()

	result, err := p.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return noResult, err
	}

	return result.LastInsertId()
}

func (p *persistence) Fetch(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := p.db.Prepare(sql)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = stmt.Close()
	}()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
