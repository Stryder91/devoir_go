package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrationDown = &cobra.Command{
	Use:   "down",
	Short: "Migration down in version for the database",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Migrating down ...")

		sqlite := mustOpenDB().DB

		sqliteDriver, err := sqlite3.WithInstance(sqlite, &sqlite3.Config{})
		if err != nil {
			fmt.Println(err)
		}

		fs, _ := (&file.File{}).Open("file://migrations")

		migration, err := migrate.NewWithInstance("install", fs, "db", sqliteDriver)
		if err != nil {
			fmt.Println(err)
		}

		version, _, _ := migration.Version()

		/* Etape 6 : migrate up - et si on est déjà à la version 2 on ne peut pas monter plus tôt.*/
		if version == 1 {
			migration.Migrate(2)
		} else if version == 0 {
			migration.Migrate(1)
		} else {
			fmt.Println("Version maximale")
		}

		if err := migration.Down(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	migrationsCmd.AddCommand(migrationDown)
}
