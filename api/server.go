package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/grahms/geolocationservice/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.GET("geolocations/:ip", server.getGeolocationByIP)
	server.router = router
	return server
}

func (server *Server) Start(addr string) error {
	return server.router.Run(addr)
}
