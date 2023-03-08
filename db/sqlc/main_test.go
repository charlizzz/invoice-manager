package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(os.Getenv("POSTGRES_DRIVER"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
