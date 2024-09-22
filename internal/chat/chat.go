package chat

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

// WebsocketHandler retorna uma func http.HandlerFunc que lida com conexoes websocket.
func WebsocketHandler(server *chatServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade da conexao HTTP para websocket.
		conn, err := websocket.Accept(w, r, nil)
		if err != nil {
			log.Println("Error upgrading connection to websocket:", err)
			return
		}

		// Cria um subscriber para a conexao.
		sub := CreateSubscriber(r.RemoteAddr)

		// Envia o subscriber para ser registrado.
		server.register <- sub

		// Goroutines para ler e escrever mensagens para subscribers conectados.
		go server.writePump(sub, conn)
		server.readPump(sub, conn)
	}
}

// writePump envia mensagens para o subscriber conectado com um timeout de 10 segundos.
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

// readPump le mensagens do subscriber conectado.
func (s *chatServer) readPump(sub *Subscriber, conn *websocket.Conn) {
	defer func() {
		s.unregister <- sub
		conn.Close(websocket.StatusNormalClosure, "closing connection")
		close(sub.send)
	}()

	for {
		_, message, err := conn.Read(context.Background())
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure || websocket.CloseStatus(err) == websocket.StatusGoingAway {
			break
		}
		if err != nil {
			log.Println("Error reading from websocket:", err)
			break
		}

		// Adicionar o identificador e a cor do usuário à mensagem.
		formattedMessage := fmt.Sprintf("<div style='color:%s'>%s</div><div style='margin-bottom: 20px'>%s</div><hr><br>", sub.color, sub.ip, string(message))
		s.broadcast <- []byte(formattedMessage)
	}
}
