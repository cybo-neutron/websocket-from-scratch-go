package websocket

import (
	"crypto/sha1"
	"encoding/base64"
)

func HashKey(key string) string {
	hash := sha1.New()
	hash.Write([]byte(key))
	hash.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))

	finalString := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return finalString
}
