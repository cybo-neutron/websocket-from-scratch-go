package websocket

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type Conn interface {
	Close() error
}

type Websocket struct {
	conn    net.Conn
	bufrw   *bufio.ReadWriter
	headers http.Header
}

func New(w http.ResponseWriter, r *http.Request) (*Websocket, error) {
	hijacker, ok := w.(http.Hijacker)

	if !ok {
		return nil, errors.New("error establishing connection")
	}

	conn, bufrw, error := hijacker.Hijack()
	if error != nil {
		return nil, errors.New("error hijacking connection")
	}
	return &Websocket{conn, bufrw, r.Header}, nil
}

func (ws *Websocket) Handshake() error {
	wsKey := ws.headers.Get("Sec-WebSocket-Key")
	fmt.Println("Sec-WebSocket-Key : ", wsKey)

	hashedKey := HashKey(wsKey)

	lines := []string{
		"HTTP/1.1 101 Switching Protocols ",
		"Upgrade: websocket ",
		"Connection: Upgrade",
		"Sec-WebSocket-Accept: " + hashedKey,
		"",
		"",
	}

	responseString := strings.Join(lines, "\r\f")

	return ws.Write(responseString)
}

func (ws *Websocket) Write(data string) error {
	_, err := ws.bufrw.Write([]byte(data))
	if err != nil {
	}

	return ws.bufrw.Flush()
}

func (ws *Websocket) read(num int) ([]byte, error) {
	buff := make([]byte, num)
	_, error := ws.bufrw.Read(buff)
	if error != nil {
		return nil, errors.New("error Reading data from socket")
	}
	return buff, nil
}
