package write

import "github.com/jmoiron/sqlx"

type Write struct {
	db *sqlx.DB
}

type WriteStmt struct {
	stmt *sqlx.Stmt
}

func New(db *sqlx.DB) *Write {
	return &Write{
		db: db,
	}
}
