package main

import (
	"fmt"
	"github.com/whenspeakteam/pg"
)

var db *pg.DB

func PgConnect() {
	var dbConfig = &pg.Options{
		User:     config["DATABASE_USER"],
		Password: config["DATABASE_PASSWORD"],
		Database: config["DATABASE_NAME"],
		Addr:     fmt.Sprintf("%s:%s", config["POSTGRES_HOST"], config["POSTGRES_PORT"]),
	}
	db = pg.Connect(dbConfig)
}
