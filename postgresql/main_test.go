package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"

	"github.com/Rafaela314/instituto-maternar/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	cfg, err := util.LoadConfig("..")
	if err != nil {
		log.Fatal("could not load config:", err)
	}

	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())

}
