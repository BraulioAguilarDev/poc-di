package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func Create(db *sqlx.DB) error {
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	err := tx.Commit()
	return err
}

func GetPerson(db *sqlx.DB) ([]Person, error) {
	people := []Person{}
	err := db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	if err != nil {
		return nil, err
	}

	return people, nil
}
