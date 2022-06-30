package api

import "github.com/gin-gonic/gin"

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
	ctx.JSON(200, location)
}
