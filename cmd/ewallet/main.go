package main

import (
	"log"

	"EWallet/internal/app"
)

func main() {
	log.Print("EWallet start...")

	a, err := app.New()
	if err != nil {
		log.Fatal("can't init app:", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal("app stop:", err)
	}
	log.Print("server stop without errors")
}
