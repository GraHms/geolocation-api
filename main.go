package main

import (
	"database/sql"
	"github.com/grahms/geolocationservice/api"
	db "github.com/grahms/geolocationservice/db/sqlc"
	"github.com/grahms/geolocationservice/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot read configurations", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddr)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
