package cmd

import (
	"github.com/spf13/cobra"
)

var v1 = &cobra.Command{
	Use:   "v1",
	Short: "Easy maker for migrations database in Go",
	Long: `You can make easy migrations for database with this 
	  		migrations CLI in GO. Documentation of Cobra https://github.com/spf13/cobra`,
}

func init() {
	rootCmd.AddCommand(v1)
}
