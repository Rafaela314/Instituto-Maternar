package main

import (
	"log"
	"os"
	"testing"

	"github.com/Rafaela314/instituto-maternar/util"
)

func TestMain(m *testing.M) {
	_, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config:", err)
	}

	os.Exit(m.Run())

}
