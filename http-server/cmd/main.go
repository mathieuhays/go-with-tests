package main

import (
	"hello/http-server"
	"log"
	"net/http"
)

func main() {
	server := http_server.NewPlayerServer(http_server.NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":42069", server))
}
