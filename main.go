package main

import (
	"database/sql"
	"fmt"
	"log"

	//_ "git.com/lib/pq"
	"github.com/Rafaela314/instituto-maternar/api"
	"github.com/Rafaela314/instituto-maternar/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}
	fmt.Println(conn.Driver())

	server := api.NewServer()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server")
	}

}
