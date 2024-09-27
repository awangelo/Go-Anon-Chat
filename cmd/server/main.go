package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/awangelo/Go-Anon-Chat/db/sqlc"
	"github.com/awangelo/Go-Anon-Chat/internal/chat"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var queries *sqlc.Queries

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	time.Sleep(10 * time.Second)

	// NewConfig deve ser usado para criar um novo Config com valores padrao.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")
	cfg.DBName = os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	queries = sqlc.New(db)
}

func main() {
	port := os.Getenv("PORT")
	server := chat.NewChatServer(queries)
	go server.Run()

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/web/static/", http.StripPrefix("/web/static/", fs))
	mux.HandleFunc("/", chat.IndexHandler)
	mux.HandleFunc("/chat", chat.WebsocketHandler(server))

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
