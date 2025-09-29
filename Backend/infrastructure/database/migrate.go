package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, directory string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: directory,
	}
	_, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Println("Succesfully Migrated database..")
	return nil

}
