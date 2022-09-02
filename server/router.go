package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nana4rider/remote-switch/controller"
)

func NewRouter() (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/docs", "docs")

	r := e.Group("/v1")

	r.GET("/computers", controller.FindAllComputers)

	r.GET("/computers/:computerId", controller.FindComputer)
	r.POST("/computers", controller.CreateComputer)
	r.PUT("/computers/:computerId", controller.UpdateComputer)
	r.DELETE("/computers/:computerId", controller.DeleteComputer)

	r.GET("/computers/:computerId/power", controller.GetPower)
	r.PUT("/computers/:computerId/power", controller.UpdatePower)

	return e, nil
}
