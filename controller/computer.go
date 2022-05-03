package controller

import (
	"context"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/nana4rider/remote-switch/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func FindComputerById(c echo.Context) (*models.Computer, error) {
	computerID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.Logger().Fatal(err)
		return nil, err
	}

	return models.FindComputer(context.Background(), boil.GetContextDB(), computerID)
}

func FindAllComputers(c echo.Context) error {
	computers, err := models.Computers().All(context.Background(), boil.GetContextDB())
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	return c.JSON(http.StatusOK, computers)
}

func FindComputer(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	return c.JSON(http.StatusOK, computer)
}

func CreateComputer(c echo.Context) error {
	computer := new(models.Computer)
	if err := c.Bind(computer); err != nil {
		return err
	}

	if errors := validateComputer(computer); errors.Has() {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if err := computer.Insert(context.Background(), boil.GetContextDB(), boil.Infer()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{"Create failed"})
	}

	if err := computer.Reload(context.Background(), boil.GetContextDB()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{"Reload failed"})
	}

	return c.JSON(http.StatusCreated, computer)
}

func UpdateComputer(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	id := computer.ID
	if err := c.Bind(computer); err != nil {
		return err
	}
	computer.ID = id

	if errors := validateComputer(computer); errors.Has() {
		return c.JSON(http.StatusBadRequest, errors)
	}

	computer.Update(context.Background(), boil.GetContextDB(), boil.Infer())

	return c.String(http.StatusNoContent, "")
}

func DeleteComputer(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	computer.Delete(context.Background(), boil.GetContextDB())

	return c.String(http.StatusNoContent, "")
}

var reIpAddr = regexp.MustCompile(`^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
var reMacAddr = regexp.MustCompile(`(?i)^[0-9a-f]{2}(:[0-9a-f]{2}){5}$`)

func validateComputer(computer *models.Computer) *ResValidationError {
	verr := new(ResValidationError)

	if len(computer.Name) == 0 {
		verr.Add("name", "Name is empty")
	}

	if !reIpAddr.MatchString(computer.IPAddress) {
		verr.Add("ip_address", "IP address format is incorrect")
	}

	if !reMacAddr.MatchString(computer.MacAddress) {
		verr.Add("mac_address", "MAC address format is incorrect")
	}

	return verr
}
