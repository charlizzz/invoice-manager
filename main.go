package main

import (
	"database/sql"
	"log"

	"github.com/charlizzz/invoice-manager/api"
	db "github.com/charlizzz/invoice-manager/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	DbDriver     = "postgres"
	DbSource     = "postgresql://root:secret@localhost:5432/invoice-db?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(DbDriver, DbSource)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
