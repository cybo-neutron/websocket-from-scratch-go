package routes

import (
	"net/http"
)

func HandleWebsocketRequest(w http.ResponseWriter, r *http.Request) {
	// 1. Handshake
	// 2. Hijack connection
	// 3. Send response to client about connection upgradation
}

func HandleWebSocketRoutes() {
	http.HandleFunc("/ws", HandleWebsocketRequest)
}
