package data

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // Postgres driver
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URL")
	return sql.Open("postgres", uri)
}
