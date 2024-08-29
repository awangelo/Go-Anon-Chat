package main

import (
	"log"
	"net/http"
	"os"

	"github.com/awangelo/Go-Anon-Chat/internal/chat"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", chat.IndexHandler)
	//mux.HandleFunc("/chat", chat.WebsocketHandler)

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
