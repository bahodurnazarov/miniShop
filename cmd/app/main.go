package main

import (
	"github.com/bahodurnazarov/miniShop/internal/server"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Since this is a test task, I decided to leave the configurations here.
	// In the real problem I would read from the environment variables.
	config := server.Config{
		BindAddr:    ":8080",
		DatabaseURL: "host=localhost port=5432 user=postgres dbname=miniShop password=postgres sslmode=disable",
	}
	log.Println(config)

	if err := server.Start(&config); err != nil {
		log.Fatal(err)
	}
}
