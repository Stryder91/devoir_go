package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrationUp = &cobra.Command{
	Use:   "up",
	Short: "Migration up in version for the database",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Migrating up ...")
		// ctx := context.Background()
		// _, err := mustOpenDB().BeginTx(ctx, model.TxOptions)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		sqlite := mustOpenDB().DB

		sqliteDriver, err := sqlite3.WithInstance(sqlite, &sqlite3.Config{})
		if err != nil {
			fmt.Println(err)
		}

		fs, _ := (&file.File{}).Open("file://migrations")

		migration, err := migrate.NewWithInstance("install", fs, "db", sqliteDriver)

		if err != nil {
			fmt.Println("err", err)
		}

		version, _, _ := migration.Version()

		/* Etape 6 : migrate down - et si on est déjà à la version 1 on le down.*/
		if version == 2 {
			migration.Migrate(1)
		} else {
			migration.Down()
			fmt.Println("Version minimale")
		}

		if err = migration.Migrate(1); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	migrationsCmd.AddCommand(migrationUp)
}
