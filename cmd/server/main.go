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

	// Inicializa o chat server
	chat.NewChatServer()
}

func main() {
	port := os.Getenv("PORT")

	server := chat.NewChatServer()
	go server.Run()

	wsHandler := chat.WebsocketHandler(server)

	mux := http.NewServeMux()
	mux.HandleFunc("/", chat.IndexHandler)
	mux.HandleFunc("/chat", wsHandler)

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
