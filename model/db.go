package model

import (
	"context"
	"database/sql"
	"fmt"
)

/*
TxOptions is the transaction isolation level used by this project.
*/
var TxOptions = &sql.TxOptions{Isolation: sql.LevelLinearizable}

/*
Truncate empties the specified table. It uses DELETE because SQLite does not
have a TRUNCATE TABLE statement, but uses a DELETE optimization instead.

See https://sqlite.org/lang_delete.html#the_truncate_optimization

The unsafeTable argument is NOT sanitized, so be sure never to pass userland data.
*/
func Truncate(ctx context.Context, tx *sql.Tx, unsafeTable string, verbose bool) error {
	res, err := tx.ExecContext(ctx, `DELETE FROM `+unsafeTable)
	if err != nil {
		return err
	}
	if verbose {
		affected, err := res.RowsAffected()
		if err != nil {
			return err
		}
		fmt.Fprintf(OutVerb, "Removed %d existing %s items\n", affected, unsafeTable)
	}
	return nil
}
