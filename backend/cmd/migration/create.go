package migration

import (
	"errors"
	"log"
	"gitlab.com/project-quiz/database"

	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

const migrationPath = "./database/migrations"

func createMigration(name string) {
	db := database.Connection()
	if err := goose.Create(db, migrationPath, name, "sql"); err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()
}

var CreateMigrationCmd = &cobra.Command{
	Use:                   "make [ARG]",
	Short:                 "Generate migration file",
	Long:                  "Generate migration file",
	DisableFlagsInUseLine: true,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		createMigration(args[0])
	},
}
