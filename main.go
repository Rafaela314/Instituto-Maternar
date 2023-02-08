package main

import (
	//"database/sql"
	//"fmt"
	"database/sql"
	"fmt"
	"log"

	_ "git.com/lib/pq"
	"github.com/Rafaela314/instituto-maternar/api"
	db "github.com/Rafaela314/instituto-maternar/db/sqlc"
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

	store := db.NewStore(conn)

	fmt.Println("\n DB CONNCTED SUCCESFULLY")

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server")
	}

}
