package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/the-arcade-01/anime-poll-app/internal/api"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("[NewAppConfig] error loading env props")
	}
}

func main() {
	server := api.NewServer()
	log.Println("[main] server running on port :8080")
	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		log.Printf("[main] error starting server, %v\n", err)
		return
	}
}
