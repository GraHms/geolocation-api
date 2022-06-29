package main

import (
	"database/sql"
	"github.com/grahms/geolocation-service/api"
	db "github.com/grahms/geolocation-service/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	DBDriver   = "postgres"
	DBSource   = "postgresql://postgres:west04@localhost:5432/geoloationsdb?sslmode=disable"
	serverAddr = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	store := db.New(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddr)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
