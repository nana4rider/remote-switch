package controller

import (
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-ping/ping"
	"github.com/labstack/echo/v4"
	"github.com/linde12/gowol"
	"golang.org/x/crypto/ssh"
)

const (
	// or "shutdown /s /t 0"
	COMMAND_POWEROFF_WINDOWS = "rundll32.exe powrprof.dll,SetSuspendState 0,1,0"
	COMMAND_POWEROFF_LINUX   = "sudo systemctl poweroff"
)

func GetState(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	ping.NewPinger(computer.IPAddress)
	pinger, err := ping.NewPinger(computer.IPAddress)
	if err != nil {
		return err
	}

	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.SetPrivileged(true)
	err = pinger.Run()

	var result ResBoolError
	if err == nil {
		result = ResBoolError{"", pinger.PacketsRecv > 0}
	} else {
		result = ResBoolError{err.Error(), false}
	}

	return c.JSON(http.StatusOK, result)
}

func SendPowerOn(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	packet, err := gowol.NewMagicPacket(computer.MacAddress)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{"The set MAC address is invalid"})
	}

	ip := net.ParseIP(computer.IPAddress).To4()
	mask := ip.DefaultMask()

	broadcast := ip
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}

	err = packet.Send(broadcast.String())
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{"Failed to send magic packet"})
	}

	return c.JSON(http.StatusOK, ResBoolError{"", true})
}

func SendPowerOff(c echo.Context) error {
	computer, err := FindComputerById(c)
	if err != nil {
		code := http.StatusNotFound
		return c.JSON(code, ResError{http.StatusText(code)})
	}

	var sshKey string
	if computer.SSHKey.Valid {
		sshKey = computer.SSHKey.String
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
		}
		sshKey = filepath.Join(homeDir, ".ssh/id_rsa")
	}

	privKey, err := ioutil.ReadFile(sshKey)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}

	signer, err := ssh.ParsePrivateKey(privKey)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}

	sconf := &ssh.ClientConfig{
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         3 * time.Second,
	}

	if computer.SSHUser.Valid {
		sconf.User = computer.SSHUser.String
	} else {
		user, err := user.Current()
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
		}
		sconf.User = user.Username
	}

	sshAddr := computer.IPAddress + ":"
	if computer.SSHPort.Valid {
		sshAddr += strconv.FormatInt(computer.SSHPort.Int64, 10)
	} else {
		sshAddr += "22"
	}

	client, err := ssh.Dial("tcp", sshAddr, sconf)
	if err != nil {
		// already off
		return c.JSON(http.StatusOK, ResBoolError{err.Error(), false})
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}
	defer session.Close()

	var command string
	if out, err := session.Output("ver"); err == nil && strings.Contains(string(out), "Microsoft Windows") {
		command = COMMAND_POWEROFF_WINDOWS
	} else {
		command = COMMAND_POWEROFF_LINUX
	}

	session, err = client.NewSession()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}
	defer session.Close()

	if err = session.Start(command); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, ResError{err.Error()})
	}

	return c.JSON(http.StatusOK, ResBoolError{"", true})
}
