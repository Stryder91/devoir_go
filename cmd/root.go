package cmd

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devoir10",
	Short: "Root command - Easy maker for migrations database in Go",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mustOpenDB() *sqlx.DB {
	sqlite3, err := sqlx.Open("sqlite3", "db")
	if err != nil {
		fmt.Println(err)
	}

	if err = sqlite3.Ping(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Database Open")

	return sqlite3
}
