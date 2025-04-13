package printer

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// PrintIntroduction welcomes the user :)
func PrintIntroduction() {
	fmt.Println("\t#########################")
	fmt.Println("\t###   H A N g M A N   ###")
	fmt.Println("\t#########################")
	fmt.Println("\t......STARTING GAME......")
	fmt.Println("")
	fmt.Println("")
}

// PrintScore is self explanatory
func PrintScore(wins, loses int) {
	ClearScreen()
	fmt.Printf("\t#########################\n")
	fmt.Printf("\t\tYOUR SCORE\n")
	fmt.Printf("\t%d wins  \t %d losses\n", wins, loses)
	fmt.Printf("\t#########################\n")
	fmt.Printf("\t        PRESS ENTER \n")
	fmt.Print("\t\t     ")
	fmt.Scanln()
	ClearScreen()
}

// PlayAgain prompts the player to decide whether they want to play again.
// It continuously asks for input until the user enters 'Y' (yes) or 'N' (no).
// Returns true if the player chooses to play again, otherwise returns false.
func PlayAgain() bool {
	for {
		fmt.Printf("\tWANT TO PLAY AGAIN ? [Y/N]\t")

		again := ""
		fmt.Scanln(&again)

		if again == "Y" || again == "y" {
			return true
		} else if again == "N" || again == "n" {
			return false
		} else {
			fmt.Println("\tEnter 'Y' or 'N'")
		}
	}
}

// ClearScreen is self explanatory
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// DrawHangman displays the Hangman figure based on the number of incorrect guesses (stage).
// The function takes two parameters:
// - stage: an integer representing the number of incorrect guesses, ranging from 0 (no mistakes) to 12 (game over).
// - guessed: a string containing the letters guessed so far.
//
// The Hangman figure is drawn progressively as the stage increases, with different body parts being added.
// The function also displays the guessed letters corresponding to each body part.
func DrawHangman(stage int, guessed string) {
	switch stage {
	case 0:
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 1:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 2:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t   |     |\t   %s\n", strings.ToUpper(string(guessed[1])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 3:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|     |\t   %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 4:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 5:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^       |\t   %s\n", strings.ToUpper(string(guessed[4])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 6:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^   ^   |\t   %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 7:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^ | ^   |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])), strings.ToUpper(string(guessed[6])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 8:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^ | ^   |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])), strings.ToUpper(string(guessed[6])))
		fmt.Printf("\t\t  /      |\t%s\n", strings.ToUpper(string(guessed[7])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 9:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^ | ^   |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])), strings.ToUpper(string(guessed[6])))
		fmt.Printf("\t\t _/      |\t%s  %s\n", strings.ToUpper(string(guessed[7])), strings.ToUpper(string(guessed[8])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 10:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^ | ^   |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])), strings.ToUpper(string(guessed[6])))
		fmt.Printf("\t\t _/ \\    |\t%s  %s  %s\n", strings.ToUpper(string(guessed[7])), strings.ToUpper(string(guessed[8])), strings.ToUpper(string(guessed[9])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 11:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   0     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t  /|\\    |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t ^ | ^   |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])), strings.ToUpper(string(guessed[6])))
		fmt.Printf("\t\t _/ \\_   |\t%s  %s  %s  %s\n", strings.ToUpper(string(guessed[7])), strings.ToUpper(string(guessed[8])), strings.ToUpper(string(guessed[9])), strings.ToUpper(string(guessed[10])))
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	case 12:
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Println("")
		fmt.Printf("\t\t   +-----+\n")
		fmt.Printf("\t\t   |     |\n")
		fmt.Printf("\t\t   |     |\t      %s\n", strings.ToUpper(string(guessed[0])))
		fmt.Printf("\t\t   |     |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[1])), strings.ToUpper(string(guessed[2])), strings.ToUpper(string(guessed[3])))
		fmt.Printf("\t\t   0     |\t   %s  %s  %s\n", strings.ToUpper(string(guessed[4])), strings.ToUpper(string(guessed[5])), strings.ToUpper(string(guessed[6])))
		fmt.Printf("\t\t  /|\\    |\t%s  %s  %s  %s  %s\n", strings.ToUpper(string(guessed[7])), strings.ToUpper(string(guessed[8])), strings.ToUpper(string(guessed[9])), strings.ToUpper(string(guessed[10])), strings.ToUpper(string(guessed[11])))
		fmt.Printf("\t\t ^ | ^   |\n")
		fmt.Printf("\t\t _/ \\_   |\n")
		fmt.Printf("\t\t         |\n")
		fmt.Printf("\t\t  R.I.P  |\n")
		fmt.Printf("\t\t===========\n\t%v/12 Guesses\n", stage)
		fmt.Println("")
	}
}
