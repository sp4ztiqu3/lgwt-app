package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/sp4ztiqu3/lgwt-app"
)

const dbFileName = "game.db.json"

func main() {
	store, clean, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf("error creating player store %v", err)
	}
	defer clean()

	fmt.Println("Let's player poker")
	fmt.Println("Type {Name} wins to record a win")

	game := poker.NewGame(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCLI(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
