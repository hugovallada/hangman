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
		fmt.Print("\n\n")
		game.printHangmanState()
		game.getGameState()
		fmt.Println(game.guessedWord)
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
		game.guess().calculateNumberOfFailedTries()
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

func (game *hangmanGame) printHangmanState() {
	file, err := os.ReadFile(fmt.Sprintf("./states/hangman%d.txt", game.numberOfFailedTries))
	if err != nil {
		game.gameOverMessage = err.Error()
		game.gameOver = true
	}
	fmt.Println(string(file))
}

func (game *hangmanGame) getGameState() {
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
}

func (game *hangmanGame) readInput() {
	r := bufio.NewReader(os.Stdin)
	input, err := r.ReadString('\n')
	if err != nil || len(input) != 1 {
		game.currentInput = rune(1)
	}
	game.currentInput = []rune(input)[0]
}

func (game *hangmanGame) guess() *hangmanGame {
	if game.guessedLetters[game.currentInput] {
		log.Println("Voce ja tentou essa letra.")
	}
	game.guessedLetters.GuessLetter(game.currentInput)
	return game
}

func (game *hangmanGame) calculateNumberOfFailedTries() {
	if !strings.Contains(game.targetWord, string(game.currentInput)) {
		game.numberOfFailedTries++
	}
}
