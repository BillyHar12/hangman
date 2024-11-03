package main

import (
	"fmt"
	"strings"
)

func displayWord(word string, guesses []rune) {
	display := ""
	for _, letter := range word {
		if contains(guesses, letter) {
			display += string(letter) + " "
		} else {
			display += "_ "
		}
	}
	fmt.Println("Current word:", strings.TrimSpace(display))
}

func contains(guesses []rune, letter rune) bool {
	for _, g := range guesses {

		if g == letter {
			return true
		}
	}
	return false
}
