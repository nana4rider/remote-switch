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

	e.GET("/computers", controller.FindAllComputers)

	e.GET("/computers/:id", controller.FindComputer)
	e.POST("/computers", controller.CreateComputer)
	e.PUT("/computers/:id", controller.UpdateComputer)
	e.DELETE("/computers/:id", controller.DeleteComputer)

	e.GET("/computers/:id/commands/state", controller.GetState)
	e.GET("/computers/:id/commands/poweron", controller.SendPowerOn)
	e.GET("/computers/:id/commands/poweroff", controller.SendPowerOff)

	return e, nil
}
