package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xiechuanj/blog/models"
)

// databasecmd is subcommand which migrate/backup/restore Blog's database.
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database subcommand migrate/backup/restore Blog's database.",
	Long:  ``,
}

// migrateDatabaseCmd is subcommand migrate Blog's database.
var migrateDatabaseCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate subcommand migrate Blog's database.",
	Long:  ``,
	Run:   migrateDatabase,
}

// backupDatabaseCmd is subcommand backup Blog's database.
var backupDatabaseCmd = &cobra.Command{
	Use:   "backup",
	Short: "backup subcommand backup Blog's database.",
	Long:  ``,
	Run:   backupDatabase,
}

// restoreDatabaseCmd is subcommand restore Blog's database.
var restoreDatabaseCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore subcommand restore Blog's database.",
	Long:  ``,
	Run:   restoreDatabase,
}

// init()
func init() {
	RootCmd.AddCommand(databaseCmd)

	databaseCmd.AddCommand(migrateDatabaseCmd)
	databaseCmd.AddCommand(backupDatabaseCmd)
	databaseCmd.AddCommand(restoreDatabaseCmd)
}

// migrateDatabase is auto-migrate database of Dockyard.
func migrateDatabase(cmd *cobra.Command, args []string) {
	models.DB.Migrate()
}

// backupDatabase is
func backupDatabase(cmd *cobra.Command, args []string) {

}

// restoreDatabase is
func restoreDatabase(cmd *cobra.Command, args []string) {

}
