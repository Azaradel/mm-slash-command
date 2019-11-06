package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Azaradel/mm-slash-command/server"
)

func main() {
	authToken := os.Getenv("AUTH_TOKEN")
	if len(authToken) == 0 {
		log.Fatal("AUTH_TOKEN environment variable is not set")
	}

	server := server.NewServer(authToken)

	if err := http.ListenAndServe(":7890", server); err != nil {
		log.Fatalf("could not start server on port 7890: %v", err)
	}
}
