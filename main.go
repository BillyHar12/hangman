package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to Hangman!")
	player := NewPlayer("Player1")
	wordList := NewWordList()
	game := NewGame(player, wordList)

	if err := game.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Thanks for playing!")
}
