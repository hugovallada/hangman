package main

import (
	"github.com/hugovallada/hangman/game"
)

func main() {
	game.
		NewGame().
		Start().
		Play().
		GameOver()
}
