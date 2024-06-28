package write

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func (w *Write) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return w.db.ExecContext(ctx, query, args)
}

func (w *Write) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return w.db.QueryContext(ctx, query, args)
}

func (w *Write) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return w.db.BeginTxx(ctx, opts)
}

func (w *Write) Prepare(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return w.db.PreparexContext(ctx, query)
}
