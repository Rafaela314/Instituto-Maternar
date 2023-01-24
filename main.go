package main

import (
	//"database/sql"
	//"fmt"
	"database/sql"
	"fmt"
	"log"

	"github.com/Rafaela314/instituto-maternar/api"
	"github.com/Rafaela314/instituto-maternar/util"
	//_ "git.com/lib/pq"
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
	fmt.Println("\n DB CONNCTED SUCCESFULLY")

	server := api.NewServer()

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server")
	}

}
