package utils

import (
	"database/sql"

	"github.com/fermyon/spin-go-sdk/pg"
	"github.com/fermyon/spin-go-sdk/variables"
)

func OpenPostgres() (*sql.DB, error) {
	addr, err := variables.Get("db_url")
	if err != nil {
		return nil, err
	}

	return pg.Open(addr), nil
}
