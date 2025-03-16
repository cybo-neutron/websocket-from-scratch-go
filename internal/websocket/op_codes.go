package websocket

var OpCodes = map[string]byte{
	"FRAME_CONTINUE": 0x0,
	"TEXT_PAYLOAD":   0x1,
	"BINARY_PAYLOAD": 0x2,
	"CLOSE":          0x8,
	"PING":           0x9,
	"PONG":           0xA,
}
