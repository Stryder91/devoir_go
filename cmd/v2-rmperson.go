package cmd

import (
	"context"
	"devoir10_tran_iyeze/model"
	"fmt"

	"github.com/spf13/cobra"
)

var rmPersonCmd = &cobra.Command{
	Use:   "rmperson",
	Short: "Easy maker for migrations database in Go",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()
		tx, err := mustOpenDB().BeginTx(ctx, model.TxOptions)
		if err != nil {
			fmt.Println(err)
		}

		statement, error := tx.Prepare("DELETE FROM people WHERE insee = ?")

		if error != nil {
			fmt.Println(error)
		}

		statement.Exec(args[0])
		tx.Commit()

	},
}

func init() {
	v2.AddCommand(rmPersonCmd)
}
