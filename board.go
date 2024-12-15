package main

import (
	"fmt"
)

type Board struct {
	cells [][]Cell
	rows  int
	cols  int
}

func NewBoard(rows, cols int) *Board {
	cells := make([][]Cell, rows)
	for i := range cells {
		cells[i] = make([]Cell, cols)
	}
	return &Board{cells: cells, rows: rows, cols: cols}
}

func (b *Board) Initialize() {
	// Initialize some cells to be alive for demonstration
	b.cells[1][2].SetAlive(true)
	b.cells[2][3].SetAlive(true)
	b.cells[3][1].SetAlive(true)
	b.cells[3][2].SetAlive(true)
	b.cells[3][3].SetAlive(true)
}

func (b *Board) Display() {
	for _, row := range b.cells {
		for _, cell := range row {
			if cell.IsAlive() {
				fmt.Print("█ ") // Alive cell
			} else {
				fmt.Print("  ") // Dead cell
			}
		}
		fmt.Println()
	}
}

func (b *Board) NextGeneration() {
	newCells := make([][]Cell, b.rows)
	for i := range newCells {
		newCells[i] = make([]Cell, b.cols)
	}

	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.cols; j++ {
			aliveNeighbors := b.CountAliveNeighbors(i, j)
			if b.cells[i][j].IsAlive() {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					newCells[i][j].SetAlive(false) // Cell dies
				} else {
					newCells[i][j].SetAlive(true) // Cell stays alive
				}
			} else {
				if aliveNeighbors == 3 {
					newCells[i][j].SetAlive(true) // Cell becomes alive
				} else {
					newCells[i][j].SetAlive(false) // Cell stays dead
				}
			}
		}
	}

	b.cells = newCells
}

func (b *Board) CountAliveNeighbors(row, col int) int {
	directions := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		newRow, newCol := row+dir.x, col+dir.y
		if newRow >= 0 && newRow < b.rows && newCol >= 0 && newCol < b.cols {
			if b.cells[newRow][newCol].IsAlive() {
				count++
			}
		}
	}
	return count
}
