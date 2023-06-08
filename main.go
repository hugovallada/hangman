package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/hugovallada/hangman/guess"
	"github.com/hugovallada/hangman/word"
)

func main() {
	targetWord, err := word.GetRandomWord()
	if err != nil {
		log.Fatal("Não foi possível buscar a palavra, tente mais tarde.", err)
	}
	fmt.Println()
	failedTries := 0
	guessedLetters := guess.GuessedLetters{}
	for {
		printHangmanState(failedTries)
		result := getGameState(targetWord, guessedLetters)
		if result == targetWord {
			fmt.Println("Parabens, voce venceu...")
			break
		}
		fmt.Println(result)
		input, err := readInput()
		if err != nil {
			log.Fatal(err)
		}
		if len(input) != 1 || !unicode.IsLetter([]rune(input)[0]) {
			log.Println("Invalid inpuut.Please use a single letter.")
			continue
		}
		letter := []rune(input)[0]
		if guessLetter(guessedLetters, letter) {
			calculateNumberOfFailedTries(targetWord, letter, &failedTries)
			if failedTries == 9 {
				printHangmanState(failedTries)
				fmt.Println("Que pena, voce perdeu... A palavra era", targetWord)
				break
			}
		} else {
			fmt.Println("Voce ja tentou essa letra, tente outra")
		}

	}
}

func getGameState(targetWord string, guessedLetters guess.GuessedLetters) string {
	result := ""
	for _, letter := range targetWord {
		if letter == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(letter)] {
			result += fmt.Sprintf("%c", letter)
		} else {
			result += "_"
		}
	}
	return result
}

func printHangmanState(hangmanState int) error {
	file, err := os.ReadFile(fmt.Sprintf("./states/hangman%d.txt", hangmanState))
	if err != nil {
		return err
	}
	fmt.Println(string(file))
	return nil
}

func guessLetter(guessedLetters guess.GuessedLetters, letter rune) bool {
	if guessedLetters[letter] {
		return false
	}
	return guessedLetters.GuessLetter(letter) == nil
}

func calculateNumberOfFailedTries(targetWord string, letter rune, failedTries *int) {
	if !strings.Contains(targetWord, string(letter)) {
		*failedTries++
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}
