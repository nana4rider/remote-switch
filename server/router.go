package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nana4rider/remote-switch/handler"
)

func NewRouter() (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/docs", "docs")

	r := e.Group("/v1")

	r.GET("/computers", handler.FindAllComputers)

	r.GET("/computers/:computerId", handler.FindComputer)
	r.POST("/computers", handler.CreateComputer)
	r.PUT("/computers/:computerId", handler.UpdateComputer)
	r.DELETE("/computers/:computerId", handler.DeleteComputer)

	r.GET("/computers/:computerId/power", handler.GetPower)
	r.PUT("/computers/:computerId/power", handler.UpdatePower)

	return e, nil
}
