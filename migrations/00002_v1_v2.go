package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upV1V2, downV1V2)
}

const v1v2Rename = `ALTER TABLE staff RENAME TO staff_old;`
const v1v2CopyStaff = `INSERT INTO staff SELECT * FROM staff_old;`
const v1v2DropStaff = `DROP TABLE staff_old;`

func upV1V2(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	const createPerson = `
CREATE TABLE people (
	insee int NOT NULL,
	given varchar(40),
	last varchar(40)
	PRIMARY KEY (insee),
);
`
	// SQLite 3 does not support adding a foreign key to an existing table, so
	// we need to recreate it, copying data.
	const createStaffV2 = `
CREATE TABLE staff (
	insee int,
	salary int,
	company int,
	people int,
	PRIMARY KEY (insee),
	FOREIGN KEY (company) REFERENCES companies(siret)
    FOREIGN KEY (people) REFERENCES people(insee)
);`

	const updateName = `
UPDATE staff
SET name = (
	SELECT given, last
	FROM people
	WHERE staff.people = people.insee
);
`
	// Create the person schema.
	err := SimpleRun(tx, []string{createPerson})
	if err != nil {
		return err
	}
	// And insert its data.
	// err = model.StorePeopleV2()
	if err != nil {
		return err
	}

	// Now integrity will work with INSEE keys.
	err = SimpleRun(tx, []string{
		v1v2Rename,
		createStaffV2,
		v1v2CopyStaff,
		v1v2DropStaff,
		updateName,
	})
	return err
}

func downV1V2(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	// SQLite 3 does not removing a foreign key from an existing table, so
	// we need to recreate it, copying data.
	const createStaffV1 = `
	CREATE TABLE staff (
		insee int(13),
		name varchar(40),
		salary int,
		staff int,
		PRIMARY KEY (insee),
		FOREIGN KEY (staff) REFERENCES companies(siret);
	);`

	const dropPeople = `DROP TABLE people;`

	err := SimpleRun(tx, []string{
		v1v2Rename,
		createStaffV1,
		v1v2CopyStaff,
		v1v2DropStaff,
		dropPeople,
	})
	return err
}
