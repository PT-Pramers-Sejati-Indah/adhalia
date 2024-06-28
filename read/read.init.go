package read

import (
	"github.com/jmoiron/sqlx"
)

type Read struct {
	db *sqlx.DB
}

type ReadStmt struct {
	stmt *sqlx.Stmt
}

func New(db *sqlx.DB) *Read {
	return &Read{
		db: db,
	}
}
