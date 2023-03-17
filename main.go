package main

import (
	"fmt"

	"di/db"
	"di/repository"

	c "di/config"

	"github.com/alecthomas/inject"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE IF NOT EXISTS person (
	first_name text,
	last_name text,
	email text
);

CREATE TABLE IF NOT EXISTS place (
	country text,
	city text NULL,
	telcode integer
)`

func run(db *sqlx.DB) {
	fmt.Println("starting application")
	db.MustExec(schema)

	if err := repository.Create(db); err != nil {
		fmt.Printf("creating error: %v\n", err.Error())
	}

	res, err := repository.GetPerson(db)
	if err != nil {
		fmt.Printf("getting error: %v\n", err.Error())
	}

	fmt.Printf("Result: %v\n", res)
}

func main() {
	inject := inject.New()
	inject.Install(
		&db.PostgresModule{
			URL: c.Config.URL,
		},
	)
	inject.Call(run)
}
