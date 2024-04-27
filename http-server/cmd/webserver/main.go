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

	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)

	server, err := poker.NewPlayerServer(store, game)
	if err != nil {
		log.Fatalf("could not create player server %v", err)
	}

	if err := http.ListenAndServe(":42069", server); err != nil {
		log.Fatalf("could not listen on port 42069 %v", err)
	}
}
