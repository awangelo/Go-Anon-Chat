package chat

type Subscriber struct {
	ip    string
	color int
	// As mensagens de cada subscriber sao enviadas atraves deste canal.
	send chan []byte
}

func getNumberOfSubscribers() int {
	return 2
}
