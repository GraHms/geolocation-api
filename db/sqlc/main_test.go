package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries Store
var testDB *sql.DB

const (
	DBDriver = "postgres"
	DBSource = "postgresql://postgres:west04@localhost:5432/geoloationsdb?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	testQueries = NewStore(testDB)
	os.Exit(m.Run())
}
