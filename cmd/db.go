package cmd

import (
	"os"
	"log"
	"github.com/spf13/cobra"
	"github.com/corentindeboisset/golang-api/app/service"
)

// Version command
func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "db:migrate",
		Short: "Upgrade the database to the latest version",
		Long:  `Upgrade the database to the latest version`,
		Run: func(cmd *cobra.Command, args []string) {
			serviceContainer, err := service.GetContainer()
			if err != nil {
				// TODO print error and exit
				return
			}

			err = serviceContainer.Migrator.LockDatabase()
			if err != nil {
				log.Printf("The lock of the database failed with the following error:\n\n    %s\n\n", err)
				os.Exit(1)
				// print error and exit
			}
			err = serviceContainer.Migrator.RunMigrations()
			if err != nil {
				log.Printf("Running the migrations failed:\n\n    %s\n\n", err)
				os.Exit(1)
				// print error and exit
			}
			err = serviceContainer.Migrator.UnlockDatabase()
			if err != nil {
				log.Printf("The unlock of the database failed with the following error:\n\n    %s\n\n", err)
				os.Exit(1)
				// print error and exit
			}

			log.Printf("Success: Database is up-to-date.")

			// print success
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
