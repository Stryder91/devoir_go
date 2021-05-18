package cmd

import (
	"github.com/spf13/cobra"
)

var v2 = &cobra.Command{
	Use:   "v2",
	Short: "Easy maker for migrations database in Go",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
}

func init() {
	rootCmd.AddCommand(v2)
}
