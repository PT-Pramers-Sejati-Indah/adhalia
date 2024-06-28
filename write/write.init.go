package write

import "github.com/jmoiron/sqlx"

type Write struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Write {
	return &Write{
		db: db,
	}
}
