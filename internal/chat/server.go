package chat

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/awangelo/Go-Anon-Chat/db/sqlc"
)

// chatServer gerencia subscribers e mensagens.
type chatServer struct {
	// subscribers armazena todos os subscribers conectados.
	// mutex abaixo eh usado quando o map precisar ser modificado.
	// O valor eh uma struct vazia, pois precisamos apenas das chaves.
	subscribers map[*Subscriber]struct{}

	// subscribersMu protege o map de subscribers.
	// Nao precisa ser retornado em NewChatServer.
	subscribersMu sync.Mutex

	// broadcast envia mensagens para todos os subscribers.
	// Como o canal eh unbuffered, as mensagens sao enviadas em uma fila FIFO.
	broadcast chan []byte

	// register solicita registrar um subscriber.
	register chan *Subscriber

	// unregister solicita remover um subscriber.
	unregister chan *Subscriber

	// conexao com a DB
	queries *sqlc.Queries
}

// NewChatServer cria um novo chat server.
func NewChatServer(queries *sqlc.Queries) *chatServer {
	return &chatServer{
		subscribers: make(map[*Subscriber]struct{}),
		broadcast:   make(chan []byte),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		queries:     queries,
	}
}

// Run inicia o chat server.
func (s *chatServer) Run() {
	for {
		select {
		case sub := <-s.register:
			s.subscribe(sub)
			s.updateUserCount()
			log.Printf("Subscriber %v connected.", sub.ip)

			s.sendAllMessages(sub)
		case sub := <-s.unregister:
			s.unsubscribe(sub)
			s.updateUserCount()
			log.Printf("Subscriber %v disconnected.", sub.ip)
		case message := <-s.broadcast:
			s.broadcastMessage(message)
			// SALVAR NA DB
		}
	}
}

// subscribe adiciona um subscriber ativo ao chat.
func (s *chatServer) subscribe(sub *Subscriber) {
	s.subscribersMu.Lock()
	defer s.subscribersMu.Unlock()
	s.subscribers[sub] = struct{}{}
}

// unsubscribe remove um subscriber ativo do chat.
func (s *chatServer) unsubscribe(sub *Subscriber) {
	s.subscribersMu.Lock()
	defer s.subscribersMu.Unlock()
	delete(s.subscribers, sub)
}

// broadcastMessage envia uma mensagem para todos os subscribers.
func (s *chatServer) broadcastMessage(msg []byte) {
	s.subscribersMu.Lock()
	defer s.subscribersMu.Unlock()
	for sub := range s.subscribers {
		sub.send <- msg
	}
}

func (s *chatServer) sendAllMessages(sub *Subscriber) {
	ctx := context.Background()

	row, err := s.queries.GetMessages(ctx)
	if err != nil {
		log.Printf("Error getting messages: %v", err)
		return
	}

	for _, parts := range row {
		color := getColorFromIP(parts.UserIp)
		formattedMessage := fmt.Sprintf("<div style='color:%s'>%s</div><div style='margin-bottom: 20px'>%s</div><hr><br>", color, parts.UserIp, parts.Content)
		sub.send <- []byte(formattedMessage)
	}
}

// updateUserCount envia o html com o numero de usuarios conectados.
func (s *chatServer) updateUserCount() {
	countMessage := fmt.Sprintf("<div id='user-count'>%d usuários conectados.</div>", s.getActiveSubscribers())
	s.broadcastMessage([]byte(countMessage))
}

// getActiveSubscribers retorna o numero de subscribers ativos.
func (s *chatServer) getActiveSubscribers() int {
	s.subscribersMu.Lock()
	defer s.subscribersMu.Unlock()
	return len(s.subscribers)
}
