package migrations

import (
	"context"
	"database/sql"
	"exp-go/internal/database"
	"exp-go/internal/models"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUser, downCreateUser)
}

func upCreateUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return database.DB_MIGRATOR.CreateTable(&models.User{})

}

func downCreateUser(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return database.DB_MIGRATOR.DropTable(&models.User{})
}
