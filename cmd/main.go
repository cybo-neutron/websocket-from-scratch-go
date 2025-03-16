package main

import (
	"log"
	"net/http"

	"github.com/cybo-neutron/websocket-from-scratch-go/routes"
)

func main() {
	routes.HandleRoutes()
	log.Fatal(http.ListenAndServe(":5005", nil))
}
