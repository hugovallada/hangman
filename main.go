package main

import (
	"fmt"
	"log"
	"unicode"

	"github.com/hugovallada/hangman/guess"
	"github.com/hugovallada/hangman/word"
)

func main() {
	targetWord, err := word.GetRandomWord()
	if err != nil {
		log.Fatal("Não foi possível buscar a palavra, tente mais tarde.", err)
	}
	fmt.Println(targetWord)
	guessedLetters := guess.GuessedLetters{}
	letter := 'i'
	guessLetter(guessedLetters, letter)
	guessLetter(guessedLetters, 'A')
	printGameState(targetWord, guessedLetters)
}

func printGameState(targetWord string, guessedLetters guess.GuessedLetters) {
	for _, letter := range targetWord {
		if letter == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(letter)] {
			fmt.Printf("%c", letter)
		} else {
			fmt.Print("_")
		}
		fmt.Print("")
	}
	fmt.Println()
}

func guessLetter(guessedLetters guess.GuessedLetters, letter rune) {
	err := guessedLetters.GuessLetter(letter)
	if err != nil {
		fmt.Println(err)
		guessLetter(guessedLetters, letter)
	}
}
