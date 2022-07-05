package api

import (
	"github.com/gin-gonic/gin"
	api "github.com/grahms/geolocationservice/api/errors"
	db "github.com/grahms/geolocationservice/db/sqlc"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"net/http"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	errResp := api.NewAPIErrorResponse()

	router.NoRoute(func(c *gin.Context) {
		errResp.SetResourceNotFound()
		c.JSON(http.StatusNotFound, errResp.Get())
	})
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		errResp.SetMethodNotAllowed()
		c.JSON(http.StatusMethodNotAllowed, errResp.Get())
	})
	router.Use(RecoveryHandler())
	// get global Monitor object
	m := ginmetrics.GetMonitor()
	//set metric path
	m.SetMetricPath("/metrics")
	m.Use(router)
	setGauge()
	router.GET("geolocations/:ip", server.getGeolocationByIP)
	server.router = router
	return server
}

func RecoveryHandler() gin.HandlerFunc {
	errResp := api.NewAPIErrorResponse()
	errResp.SetInternalServerError()
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		c.JSON(http.StatusInternalServerError, errResp.Get())
		return
	})
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
