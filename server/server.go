package server

import (
	"strconv"

	"github.com/nana4rider/remote-switch/config"
)

func StartServer() error {
	e, err := NewRouter()
	if err != nil {
		return err
	}

	conf := config.Get()

	err = e.Start(":" + strconv.Itoa(conf.Server.Port))
	if err != nil {
		e.Logger.Fatal(err)
		return err
	}

	return nil
}
