package chat

import (
	"sync"
	"time"
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

	// startTime eh o horario de inicio do chat.
	startTime time.Time
}

// NewChatServer cria um novo chat server.
func NewChatServer() *chatServer {
	return &chatServer{
		subscribers: make(map[*Subscriber]struct{}),
		broadcast:   make(chan []byte),
		register:    make(chan *Subscriber),
		unregister:  make(chan *Subscriber),
		startTime:   time.Now(),
	}
}

// Provavel implementacao do for { select }.
// func (s *chatServer) run() {

// }
