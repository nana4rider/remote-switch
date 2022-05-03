package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server   server   `json:"server" yaml:"server"`
	Database database `json:"database" yaml:"database"`
}

type server struct {
	Port int `json:"port" yaml:"port"`
}

type database struct {
	DriverName     string `json:"driverName" yaml:"driverName"`
	DataSourceName string `json:"dataSourceName" yaml:"dataSourceName"`
}

var conf *config

func Get() *config {
	return conf
}

func Load() error {
	f, err := os.Open("config/config.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(&conf)
	return err
}
