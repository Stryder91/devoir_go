package cmd

import (
	"github.com/spf13/cobra"
)

var migrationsCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Easy maker for migrations database in Go",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(migrationsCmd)
}
