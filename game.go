package main

import (
	"fmt"
)

type Game struct {
	player      *Player
	word        string
	guesses     []rune
	maxAttempts int
	attempts    int
}

func NewGame(player *Player, wordList *WordList) *Game {
	return &Game{
		player:      player,
		word:        wordList.GetRandomWord(),
		guesses:     []rune{},
		maxAttempts: 6,
		attempts:    0,
	}
}

func (g *Game) Start() error {
	for {
		displayWord(g.word, g.guesses)
		if g.attempts >= g.maxAttempts {
			fmt.Printf("Game Over! The word was: %s\n", g.word)
			break
		}

		guess, err := g.player.MakeGuess()
		if err != nil {
			return err
		}

		if !g.isAlreadyGuessed(guess) {
			g.guesses = append(g.guesses, guess)
			if !g.isCorrectGuess(guess) {
				g.attempts++
				fmt.Printf("Incorrect guess! You have %d attempts left.\n", g.maxAttempts-g.attempts)
			}
		} else {
			fmt.Println("You've already guessed that letter!")
		}

		if g.isWordGuessed() {
			fmt.Printf("Congratulations! You've guessed the word: %s\n", g.word)
			break
		}
	}
	return nil
}

func (g *Game) isAlreadyGuessed(guess rune) bool {
	for _, g := range g.guesses {
		if g == guess {
			return true
		}
	}
	return false
}

func (g *Game) isCorrectGuess(guess rune) bool {
	for _, letter := range g.word {
		if letter == guess {
			return true
		}
	}
	return false
}

func (g *Game) isWordGuessed() bool {
	for _, letter := range g.word {
		if !g.isAlreadyGuessed(letter) {
			return false
		}
	}
	return true
}
