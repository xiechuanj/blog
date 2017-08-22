package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiechuanj/blog/models"
)

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database subcommand migrate/backup/restore Blog's database.",
	Long:  ``,
}

var migrateDatabaseCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate subcommand migrate Blog's database.",
	Long:  ``,
	Run:   migrateDatabase,
}

var backupDatabaseCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup subcommand backup Blog's database.",
	Long:  ``,
	Run:   backupDatabase,
}

var restoreDatabaseCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore subcommand restore Blog's database.",
	Long:  ``,
	Run:   restoreDatabase,
}

func init() {
	RootCmd.AddCommand(databaseCmd)

	databaseCmd.AddCommand(migrateDatabaseCmd)
	databaseCmd.AddCommand(backupDatabaseCmd)
	databaseCmd.AddCommand(restoreDatabaseCmd)
}

func migrateDatabase(cmd *cobra.Command, args []string) {
	models.DB.Migrate()
}

func backupDatabase(cmd *cobra.Command, args []string) {

}

func restoreDatabase(cmd *cobra.Command, args []string) {

}
