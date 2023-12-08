package databases

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

func NewPostgresqlDBConnection() (*sql.DB, error) {
	psql, err := sql.Open("postgres", "postgres://greenlight:pa55word@localhost/tasks?sslmode=disable")

	if err != nil {
		panic(err)
	}

	migrator, err := postgres.WithInstance(psql, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	migration, err := migrate.NewWithDatabaseInstance("file://internal/infrastructure/databases/postgresql-migrations", "postgres", migrator)

	if err != nil {
		panic(err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
	return psql, nil
}
