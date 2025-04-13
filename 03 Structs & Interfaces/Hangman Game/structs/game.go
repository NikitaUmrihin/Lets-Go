package structs

import (
	"fmt"
	"hangman/printer"
	"hangman/vocab"
	"regexp"
	"strings"
)

// Game holds the game state and all data related to Hangman.
type game struct {
	Word       string
	Stage      int    // Tracks the number of incorrect guesses
	HasWon     bool   // Flag to check if the player has won
	Guess      string // Stores the last guessed letter
	Guessed    string // Stores all guessed letters
	BadGuesses string // Stores incorrect guesses
	Dashes     string // Displays the hidden word as underscores
	NewDashes  string // Tracks revealed letters in the hidden word
	Score      score  // Track players wins and losses
}

// NewGame initializes a new instance of the Game struct with default values.
func NewGame() *game {
	return &game{
		Word:       "",
		Stage:      0,
		HasWon:     false,
		Guess:      "",
		Guessed:    "",
		BadGuesses: "",
		Dashes:     "",
		NewDashes:  "",
		Score:      *NewScore(),
	}
}

// init clears the terminal screen when the program starts.
// This is executed automatically once when the package is initialized.
func init() {
	printer.ClearScreen()
}

// RevealDashes updates the dashes string to reveal correctly guessed letters.
// It takes the original word, the guessed letter, and the current state of dashes as input.
// If the guessed letter is in the word, it replaces the corresponding dashes with the letter.
func (g game) RevealDashes(word, guess, dashes string) string {
	newDashes := "\t"

	// A rune is just a fancy name for a Unicode character
	// In Go - strings are actually sequences of bytes (not characters)
	wordRunes := []rune(word) // Slice word into runes

	// Extract just the _ characters (ignoring the tab and spaces)
	dashesRunes := []rune(strings.ReplaceAll(dashes, " ", ""))[1:] // Remove '\t' and spaces

	// Safety check : check if length is the same size
	if len(wordRunes) != len(dashesRunes) {
		fmt.Println("Error: word and dashes length mismatch")
		return dashes // Return original to avoid panic
	}

	for i, r := range dashesRunes {
		// Keep already revealed letters
		if r != '_' {
			newDashes += string(r)
		} else {
			letter := string(wordRunes[i])
			// If the guessed letter is correct, reveal it
			if guess == letter {
				newDashes += strings.ToUpper(guess)
			} else {
				newDashes += "_" // Guessed letter is incorrect
			}
		}

		if i < len(dashesRunes)-1 {
			newDashes += " " // Add space between letters
		}
	}

	return newDashes
}

// Start runs the main game loop, tracking wins and losses until the player chooses to quit.
func (g game) Start() {
	again, hasWon := g.Play()
	for {
		// 	Count wins and losses
		if hasWon {
			g.Score.Wins++
		} else {
			g.Score.Loses++
		}
		printer.PrintScore(g.Score.Wins, g.Score.Loses)

		// 	Play again ?
		if again {
			again, hasWon = g.Play()
		} else {
			break
		}
	}

	fmt.Println("\tTHANKS FOR PLAYING \t:)")
}

// Play runs a single game of Hangman.
// It takes the word length as input, selects a random word, and processes user guesses.
// Returns two booleans: whether the user wants to play again and whether they won.
func (g game) Play() (bool, bool) {

	printer.PrintIntroduction()

	g.Word = vocab.GetRandomWord()

	for {
		// Draw the current state of the hangman
		printer.DrawHangman(g.Stage, g.BadGuesses)

		// If the player reached 12 wrong guesses - man is hanged
		if g.Stage == 12 {
			fmt.Printf("\tDEAD! THE WORD WAS: %s\n", strings.ToUpper(g.Word))
			// Ask if they want to play again
			return printer.PlayAgain(), false
		}

		// Display the word with hidden/revealed letters
		if g.NewDashes != "" {
			fmt.Printf(" %s\n", g.NewDashes)
		} else {
			g.Dashes = vocab.HideWord(len(g.Word))
			fmt.Printf(" %s\n", g.Dashes)
		}

		// Ask the user to guess a letter
		fmt.Printf("\n\tGUESS A LETTER: ")
		fmt.Scanln(&g.Guess)

		// Validate input: must be a single alphabetic character
		isLetter, err := regexp.MatchString("^[a-zA-Z]", g.Guess)
		if err != nil {
			printer.ClearScreen()
			fmt.Printf("\tBAD INPUT")
			fmt.Println("\tRegex match error:", err)
			panic(err)
		}

		// Convert guess to lowercase for consistency
		g.Guess = strings.ToLower(g.Guess)

		// Handle invalid guesses
		if !isLetter {
			printer.ClearScreen()
			fmt.Printf("\tNOT A LETTER, TRY AGAIN\n")
			if len(g.Guessed) == 0 && g.Stage == 0 {
				for range 5 {
					fmt.Println("")
				}
			}
		} else if len(g.Guess) > 1 {
			printer.ClearScreen()
			fmt.Printf("\tONLY 1 LETTER, TRY AGAIN\n")
			if len(g.Guessed) == 0 && g.Stage == 0 {
				for range 5 {
					fmt.Println("")
				}
			}
		} else if strings.Contains(g.Guessed, g.Guess) {
			printer.ClearScreen()
			fmt.Printf("\tYOU ALREADY GUESSED THIS LETTER, TRY AGAIN\n")

			// Correct guess !
		} else if strings.Contains(g.Word, g.Guess) {
			printer.ClearScreen()
			fmt.Printf("\tYOU FOUND A LETTER :)\n")
			if g.Stage == 0 {
				for range 5 {
					fmt.Println("")
				}
			}
			g.Guessed += g.Guess // Add to guessed letters

			// Reveal the guessed letter in the hidden word
			if g.NewDashes != "" {
				g.NewDashes = g.RevealDashes(g.Word, g.Guess, g.NewDashes)
			} else {
				g.NewDashes = g.RevealDashes(g.Word, g.Guess, g.Dashes)
			}

			// Check if the player guessed the whole word and won
			if strings.ReplaceAll(g.NewDashes, " ", "")[1:] == strings.ToUpper(g.Word) {
				g.HasWon = true
			}
			// If the player has won, display the word and a victory message
			if g.HasWon {
				printer.ClearScreen()
				wordUppercase := strings.ToUpper(g.Word)
				printer.DrawHangman(g.Stage, g.BadGuesses)
				fmt.Printf("\t")

				// Print the fully revealed word
				// move to printer
				for i, char := range wordUppercase {
					if i > 0 {
						fmt.Printf(" ")
					}
					fmt.Print(string(char))
				}
				// Print victory message
				fmt.Println("\n\t!!!! YOU WON !!!!")

				// Ask if they want to play again
				return printer.PlayAgain(), true
			}

			// Wrong guess !
		} else {
			printer.ClearScreen()
			fmt.Printf("\tWRONG GUESS :(\n")
			g.Stage++
			g.Guessed += g.Guess
			g.BadGuesses += g.Guess
		}
	}
}
