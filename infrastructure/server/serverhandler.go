package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type ServerHandler struct {
	Server *echo.Echo
	Server_r *echo.Group
}

func NewServerHandler() (*ServerHandler) {
	serverHandler := new(ServerHandler)

	serverHandler.Server = echo.New()

	//middleware
	serverHandler.Server.Use(middleware.Logger())
	serverHandler.Server.Use(middleware.Recover())
	serverHandler.Server.Use(middleware.BodyLimit("2M"))
	serverHandler.Server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 7,
	  }))

	//cors
	serverHandler.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	serverHandler.Server_r = serverHandler.Server.Group("/restricted")
	serverHandler.Server_r.Use(middleware.JWT([]byte("secret")))

	return serverHandler
}