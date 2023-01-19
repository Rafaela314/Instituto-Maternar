package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "git.com/lib/pq"
	"github.com/Rafaela314/instituto-maternar/api"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/imdb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}
	fmt.Println(conn.Driver())

	server := api.NewServer()

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("could not start server")
	}

}
