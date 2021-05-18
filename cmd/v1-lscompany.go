package cmd

import (
	"context"
	"devoir10_tran_iyeze/model"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "lscompany",
	Short: "Easy maker for migrations database in Go",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		ctx := context.Background()
		tx, err := mustOpenDB().BeginTx(ctx, model.TxOptions)
		if err != nil {
			fmt.Println(err)
		}

		data, _ := model.GetCompanies(ctx, tx)
		model.Show(os.Stdout, data)
	},
}

func init() {
	v1.AddCommand(listCmd)
}
