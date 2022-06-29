package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateGeolocation(t *testing.T) {
	args := CreateGeolocationParams{
		CountryCode: "MZ",
		CityName:    "ALASKA",
		Longitude:   "1203",
		Latitude:    "1203",
		IpAddress:   "1203",
	}
	geoLocation, err := testQueries.CreateGeolocation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, geoLocation)
	require.Equal(t, args.Longitude, geoLocation.Longitude)
	require.Equal(t, args.IpAddress, geoLocation.IpAddress)
	require.Equal(t, args.Longitude, geoLocation.Longitude)
	require.Equal(t, args.CityName, geoLocation.CityName)
	require.NotZero(t, geoLocation.ID)
}
