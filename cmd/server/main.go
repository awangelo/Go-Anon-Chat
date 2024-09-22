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
	server := chat.NewChatServer()
	go server.Run()

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", fs))
	mux.HandleFunc("/", chat.IndexHandler)
	mux.HandleFunc("/chat", chat.WebsocketHandler(server))

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
