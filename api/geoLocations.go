package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type requestParams struct {
	IP string `uri:"ip" binding:"required"`
}

func (server *Server) getGeolocationByIP(ctx *gin.Context) {
	param := new(requestParams)
	err := ctx.BindUri(param)
	if err != nil {
		return
	}
	location, err := server.store.GetGeolocation(ctx, param.IP)
	if err == sql.ErrNoRows {
		ctx.JSON(404, err)
		return
	}
	ctx.JSON(200, location)
}
