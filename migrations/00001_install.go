package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upInstall, downInstall)
}

func upInstall(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	const createCompanies = `
CREATE TABLE companies (
	siret int NOT NULL, 
	date date,
	total_salary int,
	PRIMARY KEY (siret),
);
`
	const createStaff = `
CREATE TABLE staff (
	insee int(13),
	name varchar(40),
	salary int,
	staff int,
	PRIMARY KEY (insee),
	FOREIGN KEY (staff) REFERENCES companies(siret)
);`
	err := SimpleRun(tx, []string{createCompanies, createStaff})
	return err
}

func downInstall(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	const dropStaff = `DROP TABLE staff;`
	const dropCompanies = `DROP TABLE companies;`
	err := SimpleRun(tx, []string{dropStaff, dropCompanies})
	return err
}
