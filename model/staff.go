package model

import (
	"context"
	"database/sql"

	"github.com/bojanz/currency"
)

/*
Employee models a Company employee.
*/
type Employee struct {
	Insee        int
	Name         string
	Salary       currency.Amount
	companySiret string
}

/*
MarshalYAML implements the yaml.Marshaler interface
*/
func (s Employee) MarshalYAML() (interface{}, error) {
	ready := struct {
	}{}
	return ready, nil
}

/*
UnmarshalYAML provides the yaml.Unmarshaler interface, to support decoding salary
as a currency.Amount instead of a float64.
*/
func (s *Employee) UnmarshalYAML(u func(interface{}) error) (err error) {
	// Use a local anonymous type to avoid re-entering the method.
	tmp := struct {
		// À vous...

	}{}
	err = u(&tmp)
	if err != nil {
		return err
	}
	// À vous...
	s.Salary, err = unmarshalSalary(u)
	return err
}

/*
SQLValues returns the fields on a Employee instance as ready to be passed to a SQL statement.
*/
func (s *Employee) SQLValues() []interface{} {
	p, _ := s.Salary.MarshalBinary()
	return []interface{}{
		s.Insee,
		s.Name,
		s.Salary,
		p,
	}
}

/*
Staff is a collection of Employee instances.
*/
type Staff []Employee

func storeStaffV1(ctx context.Context, tx *sql.Tx, company Company) error {
	// use MustExec, which panics on error

	//storingStaffv1 := `INSERT INTO staff (insee, name, salary, staff) VALUES (?, ?);`

	//db := sqlx.MustConnect("sqlite3", ":memory:")

	// s := Employee{ // Instanciation du struct Employe
	// 	Insee:        2578,
	// 	Name:         "Tran",
	// 	Salary:       50,
	// 	companySiret: company.SIRET,
	// }
	//result, err := db.Exec(storingStaffv1, s.Insee, s.Name, s.Salary, s.companySiret)
	//fmt.Println(result)

	return nil
}
