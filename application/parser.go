package main

import (
	"log"
	"os"

	"github.com/landeleih/ethereum-parser/application/configuration"
)

func main() {
	app := configuration.Initialize()
	if err := app.Run(); err != nil {
		log.Println("Error starting server", err)
		os.Exit(1)
	}
}
