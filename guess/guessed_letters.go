package guess

import (
	"fmt"
	"unicode"
)

type GuessedLetters map[rune]bool

func (gl GuessedLetters) GuessLetter(letter rune) error {
	if !unicode.IsLetter(letter) {
		return fmt.Errorf("insira uma letra valida")
	}
	gl[unicode.ToLower(letter)] = true
	return nil
}
