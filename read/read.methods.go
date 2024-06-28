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

func (r *Read) Prepare(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return r.db.PreparexContext(ctx, query)
}
