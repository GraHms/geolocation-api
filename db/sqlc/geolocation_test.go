package db

import (
	"context"
	"github.com/grahms/geolocationservice/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomGeolocation(t *testing.T) Geolocations {
	args := CreateGeolocationParams{
		CountryCode: util.RandStringBytes(2),
		CityName:    util.RandStringBytes(5),
		Longitude:   util.RandStringBytes(10),
		Latitude:    util.RandStringBytes(10),
		IpAddress:   util.RandStringBytes(10),
	}

	geoLocation, err := testQueries.CreateGeolocation(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, geoLocation)
	require.Equal(t, args.Longitude, geoLocation.Longitude)
	require.Equal(t, args.IpAddress, geoLocation.IpAddress)
	require.Equal(t, args.Longitude, geoLocation.Longitude)
	require.Equal(t, args.CityName, geoLocation.CityName)
	require.NotZero(t, geoLocation.ID)
	return geoLocation
}
func TestCreateGeolocation(t *testing.T) {
	createRandomGeolocation(t)
}

func TestGetGeolocation(t *testing.T) {

	args := createRandomGeolocation(t)
	geoLocation, err := testQueries.GetGeolocation(context.Background(), args.IpAddress)
	require.NoError(t, err)
	require.NotEmpty(t, geoLocation)
	require.Equal(t, args.Longitude, geoLocation.Longitude)
	require.Equal(t, args.IpAddress, geoLocation.IpAddress)
	require.Equal(t, args.Longitude, geoLocation.Longitude)
	require.Equal(t, args.CityName, geoLocation.CityName)
	require.NotZero(t, geoLocation.ID)
}
