package chat

import (
	"log"
	"net/http"

	"github.com/coder/websocket"
)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Vamos fazer um upgrade da conexao para um websocket
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading connection to websocket:", err)
	}
	// Fechar a conexao quando a funcao terminar
	defer conn.CloseNow()

	// Continuar a logica do chat aqui
	// Provavelmente colocar o select dos channels
	// e funcoes com goroutines em outro arquivo.
}
