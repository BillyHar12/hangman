package main

import (
	"fmt"
	"time"
)

func main() {
	board := NewBoard(10, 10)
	board.Initialize()

	for {
		fmt.Print("\033[H\033[2J") // Clear the console
		board.Display()
		board.NextGeneration()
		time.Sleep(500 * time.Millisecond) // Pause for half a second
	}
}
