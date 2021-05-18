package model

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/bojanz/currency"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
)

/*
Company models a company entry.
*/
type Company struct {
	SIRET       string          `yaml:"siret"`
	Changed     time.Time       `yaml:"changed"`
	TotalSalary currency.Amount `yaml:"-"`
	Staff       []Employee      `yaml:"staff"`
}

/*
MarshalYAML implements the yaml.Marshaler interface
*/
func (c Company) MarshalYAML() (interface{}, error) {
	ready := struct {
		// À vous...
		SIRET       string
		Changed     time.Time
		TotalSalary currency.Amount
	}{
		// À vous...
		SIRET:       c.SIRET,
		Changed:     c.Changed,
		TotalSalary: c.TotalSalary,
	}
	return ready, nil
}

/*
UnmarshalYAML provides the yaml.Unmarshaler interface, to support decoding salary
as a currency.Amount instead of a float64.
*/
func (c *Company) UnmarshalYAML(u func(interface{}) error) (err error) {
	// Use a local anonymous type to avoid re-entering the method.
	tmp := struct {
		// À vous...
		SIRET       string
		Changed     time.Time
		TotalSalary float64
	}{}
	err = u(&tmp)
	if err != nil {
		return err
	}
	// À vous...
	c.SIRET = tmp.SIRET
	c.Changed = tmp.Changed
	c.TotalSalary, err = unmarshalTotalSalary(u)
	return
}

/*
SQLValues returns the fields on a Employee instance as ready to be passed to a SQL statement.
*/
func (c *Company) SQLValues() []interface{} {
	p, _ := c.TotalSalary.MarshalBinary()
	fmt.Println(p)
	return []interface{}{
		c.SIRET,
		c.Changed,
		p,
	}

}

/*
Companies is a collection of Company instances.
*/
type Companies []Company

// Comes from a LEFT OUTER JOIN, so all fields are nullable.
type dbStaff struct {
	// À vous...
}
type dbCompany struct {
	siret        string
	date         time.Time
	total_salary currency.Amount
}

/*
StoreCompanyV1 stores a Company instance in the database in V1 schema.
*/
func StoreCompanyV1(ctx context.Context, tx *sql.Tx, company Company) error {

	storingCompany := `INSERT INTO companies (siret, date, total_salary) VALUES (?, ?, ?);`

	query, err := tx.PrepareContext(ctx, storingCompany)
	if err != nil {
		fmt.Println(err)
	}

	_, err = query.ExecContext(ctx, company.SIRET, company.Changed, company.TotalSalary)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

/*
GetCompanyWithoutStaff retrieves a single company instance by its id, without its staff.
*/
func GetCompanyWithoutStaff(ctx context.Context, tx *sql.Tx, siret string) (Company, error) {
	// À vous... pensez à tx.QueryRowContext
	db := sqlx.MustConnect("sqlite3", ":memory:")
	rows, err := db.Query("SELECT siret, date, total_salary FROM company WHERE siret=" + siret + ";")
	fmt.Println(rows)
	myCompany := Company{
		SIRET: siret,
	}
	return myCompany, err
}

/*
GetCompanies retrieves all Company instances from the database in V1 schema.
*/
func GetCompanies(ctx context.Context, tx *sql.Tx) (Companies, error) {
	// Usually don't loop on SQL queries: use a join in a single query instead.
	// À vous...
	var dbCompanies []dbCompany
	tmpCompany := dbCompany{}
	query, err := tx.QueryContext(ctx, "select siret,date,total_salary from companies")
	if err != nil {
		fmt.Println(err)
	}

	for query.Next() {
		if err = query.Scan(&tmpCompany.siret, &tmpCompany.date, &tmpCompany.total_salary); err != nil {
			fmt.Println(err)
		}
		dbCompanies = append(dbCompanies, tmpCompany)
	}

	return companiesFromDBCompanies(dbCompanies), nil
}

func staffFromDBStaff(dbst []dbStaff) Staff {
	rows := make(Staff, 0, len(dbst))
	// À vous...
	return rows
}

func companiesFromDBCompanies(dbcs []dbCompany) Companies {
	companies := make(Companies, 0, len(dbcs))
	// À vous...
	for _, cs := range dbcs {
		c := Company{}
		c.SIRET = cs.siret
		c.Changed = cs.date
		c.TotalSalary = cs.total_salary
		companies = append(companies, c)
	}
	return companies
}

/*
Show outputs a text representation of a Companies slice, using YAML as the actual format.
*/
func Show(w io.Writer, companies Companies) {
	// À vous...
	y, _ := yaml.Marshal(companies)
	fmt.Fprintf(w, "%s\n", string(y))
}
