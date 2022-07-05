package usecase

import (
	"context"
	"database/sql"
	"fmt"
	db "github.com/grahms/geolocationservice/db/sqlc"
	"github.com/grahms/geolocationservice/lib"
)

var csv = lib.NewCsv()

type Csv2Db struct {
	store    db.Store
	filepath string
}

func New(store db.Store, filepath string) *Csv2Db {
	return &Csv2Db{store: store,
		filepath: filepath,
	}
}

func (u *Csv2Db) AddLocationToDb() {
	csv.Parse(u.filepath, u.readLocationsFromChan)
}

func (u *Csv2Db) readLocationsFromChan(loc lib.GeoCSV) {

	if csv.ValidateRow(&loc) {
		u.persist(&loc)
		return
	}

}

// persist locations from CSV to database
func (u *Csv2Db) persist(loc *lib.GeoCSV) {
	args := db.CreateGeolocationParams{
		CountryCode: loc.CountryCode,
		CityName:    loc.City,
		Longitude:   loc.Longitude,
		Latitude:    loc.Latitude,
		IpAddress:   loc.IPAddress,
	}
	_, err := u.store.CreateGeolocation(context.Background(), args)

	fmt.Println()
	if err != nil {
		return
	}
	if err == sql.ErrNoRows {
		return
	}
}
