package vocab

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const MIN_WORD_LENGTH = 6
const MAX_WORD_LENGTH = 12

// GetRandomWord selects a random word from a file .
// It takes a word length between 6 and 12 (inclusive) as input.
// The function reads from a file containing words of a specific length
// named "<wordLength>.txt", located in the "words" directory.
func GetRandomWord() string {
	wordLength := rand.Intn(MAX_WORD_LENGTH-MIN_WORD_LENGTH+1) + MIN_WORD_LENGTH
	filePath := "words/" + strconv.Itoa(wordLength) + ".txt"
	var fileContent []byte
	var err error

	// Only read from the file if the word length is between 6 and 12
	if wordLength >= MIN_WORD_LENGTH && wordLength <= MAX_WORD_LENGTH {
		fileContent, err = os.ReadFile(filePath)
	}
	// If an error occurs while reading the file, terminate the program
	if err != nil {
		panic(err)
	}
	// Split file content into a list of words
	words := strings.Split(string(fileContent), " ")

	// Select a random word from the list
	randomNum := rand.Intn(len(words))
	return words[randomNum-1]
}

// HideWord generates a string of underscores representing a hidden word of a given length.
// Each underscore is separated by a space to maintain readability.
func HideWord(wordLength int) string {
	dashes := "\t"
	for i := 0; i < wordLength; i++ {
		if i > 0 {
			dashes += " " // Add a space between dashes
		}
		dashes += "_"
	}
	return dashes
}
