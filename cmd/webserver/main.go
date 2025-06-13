package main

import (
	"log"
	"net/http"

	poker "github.com/sp4ztiqu3/lgwt-app"
)

const dbFileName = "game.db.json"

func main() {
	store, clean, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("error creating player store %v", err)
	}
	defer clean()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
