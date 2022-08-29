package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nana4rider/remote-switch/config"
	"github.com/nana4rider/remote-switch/server"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	conf := config.Get()

	db, err := sql.Open(conf.Database.DriverName, conf.Database.DataSourceName)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	boil.SetDB(db)

	if err := server.StartServer(); err != nil {
		panic(err)
	}
}
