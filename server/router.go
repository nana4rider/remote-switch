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

	r := e.Group("/v1")

	r.GET("/computers", controller.FindAllComputers)

	r.GET("/computers/:id", controller.FindComputer)
	r.POST("/computers", controller.CreateComputer)
	r.PUT("/computers/:id", controller.UpdateComputer)
	r.DELETE("/computers/:id", controller.DeleteComputer)

	r.GET("/computers/:id/power", controller.GetPower)
	r.PUT("/computers/:id/power", controller.UpdatePower)

	return e, nil
}
