package adhalia

import (
	"github.com/PT-Pramers-Sejati-Indah/adhalia/read"
	"github.com/PT-Pramers-Sejati-Indah/adhalia/write"
	"github.com/jmoiron/sqlx"
)

// CQRS
type Adhalia struct {
	Reader *read.Read
	Writer *write.Write
}

func NewDatabase(
	readerDriverName string,
	readerDataSourceName string,
	writerDriverName string,
	writerDataSourceName string,
) (*Adhalia, error) {
	readerDb, err := sqlx.Open(readerDriverName, readerDataSourceName)
	if err != nil {
		return nil, err
	}
	writerDb, err := sqlx.Open(writerDriverName, writerDataSourceName)
	if err != nil {
		return nil, err
	}
	return &Adhalia{
		Reader: read.New(readerDb),
		Writer: write.New(writerDb),
	}, nil
}
