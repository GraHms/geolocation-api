package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	api "github.com/grahms/geolocationservice/api/errors"
	"net/http"
)

type requestParams struct {
	IP string `uri:"ip" binding:"required"`
}

func (server *Server) getGeolocationByIP(ctx *gin.Context) {
	errResponse := api.NewAPIErrorResponse()
	param := new(requestParams)
	err := ctx.BindUri(param)
	if err != nil {
		return
	}
	location, err := server.store.GetGeolocation(ctx, param.IP)

	if err == sql.ErrNoRows {
		errResponse.SetGeoLocationNotFound()
		ctx.JSON(http.StatusNotFound, errResponse.Get())
		return
	}
	coords := Coordinates{
		Latitude:  &location.Latitude,
		Longitude: &location.Longitude}

	response := Model{
		CountryCode: &location.CountryCode,
		CityName:    &location.CityName,
		IpAddress:   &location.IpAddress,
		Coordinates: &coords,
	}

	ctx.JSON(http.StatusOK, response)
}
