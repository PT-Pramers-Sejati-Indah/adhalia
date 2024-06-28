package write

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

// do not use this method if you do not know what you're doing
func (w *Write) GetDB() *sqlx.DB {
	return w.db
}

func (w *Write) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return w.db.ExecContext(ctx, query, args)
}

func (w *Write) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return w.db.QueryContext(ctx, query, args)
}

func (w *Write) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return w.db.BeginTxx(ctx, opts)
}

func (w *Write) Prepare(ctx context.Context, query string) (*WriteStmt, error) {
	stmt, err := w.db.PreparexContext(ctx, query)
	if err != nil {
		return nil, err
	}
	return &WriteStmt{
		stmt: stmt,
	}, nil
}

func (w *WriteStmt) Exec(ctx context.Context, args ...any) (sql.Result, error) {
	return w.stmt.ExecContext(ctx, args)
}

func (w *WriteStmt) Query(ctx context.Context, args ...any) (*sql.Rows, error) {
	return w.stmt.QueryContext(ctx, args)
}
