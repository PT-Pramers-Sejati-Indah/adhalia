package read

import (
	"github.com/jmoiron/sqlx"
)

type Read struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Read {
	return &Read{
		db: db,
	}
}
