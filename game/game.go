package game

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

type hangmanGame struct {
	targetWord          string
	numberOfFailedTries int8
	guessedLetters      guess.GuessedLetters
	gameOver            bool
	gameOverMessage     string
	guessedWord         string
	currentInput        rune
}

func NewGame() *hangmanGame {
	return &hangmanGame{}
}

func (game *hangmanGame) Start() *hangmanGame {
	word, err := word.GetRandomWord()
	if err != nil {
		game.gameOverMessage = "Nao foi possivel iniciar a palavra. Tente novamente mais tarde"
		game.gameOver = true
	}
	game.targetWord = word
	game.guessedLetters = guess.GuessedLetters{}
	return game
}

func (game *hangmanGame) Play() *hangmanGame {
	for !game.gameOver {
		game.printHangmanState().
			getGameState().
			printGuessedWord()
		if game.guessedWord == game.targetWord {
			game.gameOverMessage = "Parabens, voce venceu!"
			game.gameOver = true
			break
		}
		game.readInput()
		if !unicode.IsLetter(game.currentInput) {
			log.Println("Invalid input, please use a single letter.")
			continue
		}
		game.guess()
		if game.numberOfFailedTries >= 9 {
			game.printHangmanState()
			game.gameOverMessage = fmt.Sprintf("Que pena, voce perdeu... A palavra era %s", game.targetWord)
			game.gameOver = true
			break
		}
	}
	return game
}

func (game *hangmanGame) GameOver() {
	if game.gameOver {
		fmt.Println(game.gameOverMessage)
	}
}

func (game *hangmanGame) printHangmanState() *hangmanGame {
	file, err := os.ReadFile(fmt.Sprintf("./states/hangman%d.txt", game.numberOfFailedTries))
	if err != nil {
		game.gameOverMessage = err.Error()
		game.gameOver = true
	}
	fmt.Print("\n\n")
	fmt.Println(string(file))
	return game
}

func (game *hangmanGame) getGameState() *hangmanGame {
	result := ""
	for _, letter := range game.targetWord {
		if letter == ' ' {
			result += " "
		} else if game.guessedLetters[unicode.ToLower(letter)] {
			result += fmt.Sprintf("%c", letter)
		} else {
			result += "_"
		}
	}
	game.guessedWord = result
	return game
}

func (game *hangmanGame) printGuessedWord() {
	fmt.Println(game.guessedWord)
}

func (game *hangmanGame) readInput() *hangmanGame {
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil || len(input) != 1 {
		game.currentInput = rune(1)
	}
	game.currentInput = []rune(input)[0]
	return game
}

func (game *hangmanGame) guess() {
	if game.guessedLetters[game.currentInput] {
		log.Println("Voce ja tentou essa letra.")
		return
	}
	game.guessedLetters.GuessLetter(game.currentInput)
	game.calculateNumberOfFailedTries()
}

func (game *hangmanGame) calculateNumberOfFailedTries() {
	if !strings.Contains(game.targetWord, string(game.currentInput)) {
		game.numberOfFailedTries++
	}
}
