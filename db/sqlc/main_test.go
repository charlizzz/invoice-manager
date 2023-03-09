package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	DbDriver = "postgres"
	DbSource = "postgresql://root:secret@localhost:5432/invoice-db?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(DbDriver, DbSource)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
