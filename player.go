package main

import (
	"fmt"
	"strings"
)

type Player struct {
	name string
}

func NewPlayer(name string) *Player {
	return &Player{name: name}
}

func (p *Player) MakeGuess() (rune, error) {
	var input string
	fmt.Printf("%s, enter your guess (a single letter): ", p.name)
	fmt.Scanln(&input)

	if len(input) != 1 {
		return 0, fmt.Errorf("invalid input, please enter a single letter")
	}

	return rune(strings.ToLower(input)[0]), nil
}
