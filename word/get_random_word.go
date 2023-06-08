package word

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRandomWord() (string, error) {
	request, err := http.NewRequest("GET", "https://random-word-api.herokuapp.com/word", nil)
	if err != nil {
		return "", err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var hangmanWord HangmanWord
	err = json.NewDecoder(response.Body).Decode(&hangmanWord)
	if err != nil {
		return "", err
	}
	return hangmanWordFromArray(hangmanWord)
}

func hangmanWordFromArray(array HangmanWord) (string, error) {
	if len(array) < 1 {
		return "", fmt.Errorf("nenhuma palavra encontrada")
	}
	return array[0], nil
}
