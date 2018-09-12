package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	//"net/http"
)

var (
	Server *echo.Echo
	Server_r *echo.Group
)

//init echo web server
func Init() {
	Server = echo.New()

	//middleware
	Server.Use(middleware.Logger())
	Server.Use(middleware.Recover())
	Server.Use(middleware.BodyLimit("2M"))
	Server.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 7,
	  }))

	//cors
	Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	Server_r = Server.Group("/restricted")
	Server_r.Use(middleware.JWT([]byte("secret")))
}
