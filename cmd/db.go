package cmd

import (
	"github.com/spf13/cobra"
)

// Version command
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "db:migrate",
		Short: "Upgrade the database to the latest version",
		Long:  `Upgrade the database to the latest version`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "db:exec",
		Short: "Execute (or revert) a migration",
		Long:  `Execute (or revert) a migration`,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO
		},
	})
}
