package main

import (
	"hello/http-server"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := http_server.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := http_server.NewPlayerServer(store)

	if err := http.ListenAndServe(":42069", server); err != nil {
		log.Fatalf("could not listen on port 42069 %v", err)
	}
}
