package main

import (
	"fmt"

	"github.com/Rafaela314/instituto-maternar/internal/host"
)

func main() {

	host := host.NewHost()

	fmt.Println(host.AppConfig)
}
