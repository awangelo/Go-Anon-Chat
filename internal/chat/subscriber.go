package chat

import "fmt"

type Subscriber struct {
	ip    string
	color int
	// As mensagens de cada subscriber sao enviadas atraves deste canal.
	send chan []byte
}

// Apenas para ver o para de *Subscribers
func (s *Subscriber) String() string {
	return fmt.Sprintf("ip: %v, color: %v", s.ip, s.color)
}
