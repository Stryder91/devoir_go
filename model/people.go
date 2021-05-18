package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

/*
Person models a Person used in Employee instances within Company instances.
*/
type Person struct {
	given  string
	last   string
	insee  int
	salary int
}

/*
SQLValues returns the fields on a Person instance as ready to be passed to a SQL statement.
*/
func (p *Person) SQLValues() []interface{} {
	return []interface{}{}
}

/*
People is a collection of Person.
*/
type People []Person

/*
GetPerson retrieves a single Person instance by its INSEE ID.
*/
func GetPerson(ctx context.Context, tx *sql.Tx, insee string) (Person, error) {
	person := Person{}

	// panics on error
	db := sqlx.MustConnect("sqlite3", ":memory:")

	schema := `SELECT * FROM people where insee=` + insee + `;`

	result, err := db.Exec(schema)

	fmt.Println(result)

	return person, err
}

/*
StorePeopleV2 stores a People instance in the v2 database.
*/
func StorePeopleV2(ctx context.Context, tx *sql.Tx, people People) error {
	// use MustExec, which panics on error
	storingPeople := `INSERT INTO people (given, last, insee, salary) VALUES (?, ?);`

	db := sqlx.MustConnect("sqlite3", ":memory:")

	result, err := db.Exec(storingPeople, "Lionel", "Tran", 2572, 500)
	fmt.Println(result)

	return err
}
