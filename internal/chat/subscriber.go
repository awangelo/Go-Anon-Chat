package chat

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

type Subscriber struct {
	ip    string
	color string
	// As mensagens de cada subscriber sao enviadas atraves deste canal.
	send chan []byte
}

// createSubscriber cria um novo subscriber com base no endereco IP.
func CreateSubscriber(remoteAddr string) *Subscriber {
	identifier, userColor := getUserIdentity(remoteAddr)
	return &Subscriber{
		ip:    identifier,
		color: userColor,
		send:  make(chan []byte, 256),
	}
}

// getUserIdentity retorna um identificador e uma cor unica para um endereco IP.
func getUserIdentity(remoteAddr string) (string, string) {
	ip, _, _ := net.SplitHostPort(remoteAddr)

	octetos := strings.Split(ip, ".")
	identifier := fmt.Sprintf(".%s..%s", octetos[1], octetos[3])

	color := getColorFromIP(ip)

	return identifier, color
}

// getColorFromIP retorna uma cor hexadecimal unica com base no endereco IP.
func getColorFromIP(ip string) string {
	hash := md5.Sum([]byte(ip))
	return "#" + hex.EncodeToString(hash[:3])
}
