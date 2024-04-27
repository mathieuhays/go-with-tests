package main

import (
	poker "hello/http-server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	store, closeStore, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeStore()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":42069", server); err != nil {
		log.Fatalf("could not listen on port 42069 %v", err)
	}
}
