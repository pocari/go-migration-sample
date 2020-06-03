package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

func NewDB() (*sql.DB, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}

	return newDB(
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		dbPort,
	)
}

func newDB(database, user, pass, host string, port int) (*sql.DB, error) {
	return sql.Open(
		"mysql",
		generateConnectionString(database, user, pass, host, port),
	)
}

func generateConnectionString(database, user, pass, host string, port int) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4",
		user,
		pass,
		host,
		port,
		database,
	)
}
