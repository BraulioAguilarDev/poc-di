package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresModule struct {
	URL string
}

func (p *PostgresModule) ProvidePostgresDB() (*sqlx.DB, error) {
	// postgres://postgres:password@localhost:5432/dbtest?sslmode=disable
	db, err := sqlx.Open("postgres", p.URL)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
