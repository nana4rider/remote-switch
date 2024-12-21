package handler

import (
	"context"
	"errors"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nana4rider/remote-switch/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	reIpAddr     = regexp.MustCompile(`^(([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([1-9]?[0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	reMacAddr    = regexp.MustCompile(`(?i)^[0-9a-f]{2}(:[0-9a-f]{2}){5}$`)
	reArpMacAddr = regexp.MustCompile(`(?i)([0-9a-f]{2}(?:[:-][0-9a-f]{2}){5})`)
)

func FindComputerById(c echo.Context) (*models.Computer, error) {
	computerID, err := strconv.Atoi(c.Param("computerId"))
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	return models.FindComputer(context.Background(), boil.GetContextDB(), computerID)
}

// @Summary コンピュータの一覧を取得
// @Tags computer
// @Produce json
// @Success 200 {array} models.Computer
// @Failure 404 {object} ResError
// @Router /computers [get]
func FindAllComputers(c echo.Context) error {
	computers, err := models.Computers().All(context.Background(), boil.GetContextDB())
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	return c.JSON(http.StatusOK, computers)
}

// @Summary コンピュータの詳細を取得
// @Tags computer
// @Produce json
// @Param computerId path int true "Computer ID"
// @Success 200 {object} models.Computer
// @Failure 404 {object} ResError
// @Router /computers/{computerId} [get]
func FindComputer(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	return c.JSON(http.StatusOK, computer)
}

// @Summary コンピュータを登録
// @Tags computer
// @Accept json
// @Param request body models.Computer true "ssh_key: default $HOME/.ssh/id_rsa, ssh_port: default 22, mac_address: default 'arp -a ip_address'"
// @Success 201
// @Failure 400 {object} ResValidationError
// @Router /computers [post]
func CreateComputer(c echo.Context) error {
	computer := new(models.Computer)
	if err := c.Bind(computer); err != nil {
		return err
	}

	if len(computer.MacAddress) == 0 && len(computer.IPAddress) != 0 {
		if mac, err := getMacAddr(computer.IPAddress); err == nil {
			computer.MacAddress = mac
		}
	}

	if errors := validateComputer(computer); errors.Has() {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if err := computer.Insert(context.Background(), boil.GetContextDB(), boil.Infer()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}

	if err := computer.Reload(context.Background(), boil.GetContextDB()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{"Reload failed"})
	}

	return c.JSON(http.StatusCreated, computer)
}

// @Summary コンピュータを更新
// @Tags computer
// @Accept json
// @Param computerId path int true "Computer ID"
// @Param request body models.Computer true "ssh_key: default $HOME/.ssh/id_rsa, ssh_port: default 22, mac_address: default 'arp -a ip_address'"
// @Success 204
// @Failure 400 {object} ResValidationError
// @Router /computers/{computerId} [put]
func UpdateComputer(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	id := computer.ComputerID
	if err := c.Bind(computer); err != nil {
		return err
	}
	computer.ComputerID = id

	if len(computer.MacAddress) == 0 && len(computer.IPAddress) != 0 {
		if mac, err := getMacAddr(computer.IPAddress); err == nil {
			computer.MacAddress = mac
		}
	}

	if errors := validateComputer(computer); errors.Has() {
		return c.JSON(http.StatusBadRequest, errors)
	}

	if _, err := computer.Update(context.Background(), boil.GetContextDB(), boil.Infer()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}

	return c.String(http.StatusNoContent, "")
}

// @Summary コンピュータを削除
// @Tags computer
// @Param computerId path int true "Computer ID"
// @Success 204
// @Failure 404 {object} ResError
// @Router /computers/{computerId} [delete]
func DeleteComputer(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	if _, err := computer.Delete(context.Background(), boil.GetContextDB()); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}

	return c.String(http.StatusNoContent, "")
}

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

func getMacAddr(ipAddr string) (string, error) {
	if out, err := exec.Command("arp", "-a", ipAddr).Output(); err == nil {
		group := reArpMacAddr.FindSubmatch(out)
		if len(group) > 0 {
			mac := strings.ReplaceAll(string(group[1]), "-", ":")
			return mac, nil
		}
	}

	return "", errors.New("failed to get")
}
