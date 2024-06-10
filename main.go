package main

import (
	"log"

	"github.com/ashishjuyal/banking/app"
	"github.com/ashishjuyal/banking/logger"
)

func main() {
	log.Println("Startin app")
	logger.Info("Starting application")
	app.Start()
}
