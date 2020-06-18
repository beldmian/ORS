package main

import (
	"github.com/beldmian/ORS/config"
	"github.com/beldmian/ORS/internal/app/db"
	"github.com/beldmian/ORS/internal/app/server"
)

func main() {
	serverConfig, err := config.NewConfig("./config/server.toml")
	if err != nil {
		panic(err)
	}
	db := db.Datebase{}
	server := server.New(serverConfig, db)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
