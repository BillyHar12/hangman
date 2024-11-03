package main

import (
	"math/rand"
	"time"
)

type WordList struct {
	words []string
}

func NewWordList() *WordList {
	return &WordList{
		words: []string{"hangman", "golang", "programming", "developer", "interface"},
	}
}

func (wl *WordList) GetRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return wl.words[rand.Intn(len(wl.words))]
}
