package chat

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

func WebsocketHandler(server *chatServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade da conexao HTTP para websocket.
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Println("Error upgrading connection to websocket:", err)
			return
		}

		// Novo subscriber que se conectou.
		sub := &Subscriber{
			ip:    r.RemoteAddr,
			color: 0, // Defina a cor como desejar
			send:  make(chan []byte, 256),
		}

		// Envia o subscriber para ser registrado.
		server.register <- sub

		// Goroutines para ler e escrever mensagens para subscribers conectados.
		go server.writePump(sub, conn)
		server.readPump(sub, conn)
	}
}

func (s *chatServer) writePump(sub *Subscriber, conn *websocket.Conn) {
	defer func() {
		conn.Close(websocket.StatusNormalClosure, "channel closed")
	}()

	for message := range sub.send {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := conn.Write(ctx, websocket.MessageText, message); err != nil {
			return
		}
	}
}

func (s *chatServer) readPump(sub *Subscriber, conn *websocket.Conn) {
	defer func() {
		s.unregister <- sub
		conn.Close(websocket.StatusNormalClosure, "closing connection")
		close(sub.send)
	}()

	for {
		_, message, err := conn.Read(context.Background())
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			break
		}
		if err != nil {
			log.Println("Error reading from websocket:", err)
			break
		}
		s.broadcast <- message
	}
}
