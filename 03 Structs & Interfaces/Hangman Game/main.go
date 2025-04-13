package main

import (
	"hangman/printer"
	"hangman/structs"
)

func init() {
	printer.ClearScreen()
}

// This Go script is a Hangman game, where the user guesses letters to reveal a hidden word.
// The program selects a random word from predefined text files,
// allows users to guess letters, and visually represents the Hangman stages.

func main() {
	g := structs.NewGame()
	g.Start()

}
