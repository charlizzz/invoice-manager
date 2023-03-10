package main

import (
	"database/sql"
	"log"

	"github.com/charlizzz/invoice-manager/api"
	db "github.com/charlizzz/invoice-manager/db/sqlc"
	"github.com/charlizzz/invoice-manager/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load the config: ", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to the database: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}
