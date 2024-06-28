package read

import (
	"context"
	"github.com/jmoiron/sqlx"
)

func (r *Read) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.GetContext(ctx, dest, query, args)
}

func (r *Read) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.SelectContext(ctx, dest, query, args)
}

func (r *Read) Prepare(ctx context.Context, query string) (*ReadStmt, error) {
	stmt, err := r.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &ReadStmt{
		stmt: stmt,
	}, nil
}

// do not use this method if you do not know how to use it
func (r *Read) GetDB() *sqlx.DB {
	return r.db
}

func (r *ReadStmt) Get(ctx context.Context, dest interface{}, args ...interface{}) error {
	return r.stmt.GetContext(ctx, dest, args)
}

func (r *ReadStmt) Select(ctx context.Context, dest interface{}, args ...interface{}) error {
	return r.stmt.SelectContext(ctx, dest, args)
}
