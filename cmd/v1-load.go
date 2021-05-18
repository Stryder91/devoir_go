package cmd

import (
	"context"
	"devoir10_tran_iyeze/model"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var loadCmd = &cobra.Command{
	Use:   "load",
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

		var companiesFromFile model.Companies

		file, err := ioutil.ReadFile("data/companies.yml")
		if err != nil {
			fmt.Println(err)
		}

		yaml.Unmarshal(file, &companiesFromFile)

		for _, company := range companiesFromFile {
			model.StoreCompanyV1(ctx, tx, company)
		}

		tx.Commit()

		fmt.Println("données ajoutées")

	},
}

func init() {
	v1.AddCommand(loadCmd)
}
