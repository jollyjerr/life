package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var deltas = [3]int{-1, 0, 1}

type board [][]bool

func newBoard(size int) board {
	cells := make(board, size)
	for i := range cells {
		cells[i] = make([]bool, size)
	}
	return cells
}

func copyBoard(b board) board {
	out := newBoard(len(b))

	for i := range b {
		for j, val := range b[i] {
			out[i][j] = val
		}
	}

	return out
}

func tick(b board) {
	tmp := copyBoard(b)

	for y := range b {
		for x := range b[y] {
			b[y][x] = fate(tmp, x, y)
		}
	}
}

func fate(b board, x, y int) bool {
	count := countLivingNeighbors(b, x, y)
	born := !(b[y][x]) && count == 3
	survive := (b[y][x]) && (count == 2 || count == 3)
	return born || survive
}

func countLivingNeighbors(b board, x, y int) int {
	size := len(b)

	count := 0
	for _, i := range deltas {
		for _, j := range deltas {
			if i == 0 && j == 0 {
				continue
			}

			nY := (y + j + size) % size
			nX := (x + i + size) % size

			if b[nY][nX] {
				count++
			}
		}
	}

	return count
}

func render(b board) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for _, row := range b {
		for _, col := range row {
			if col {
				fmt.Print("*", " ")
			} else {
				fmt.Print(".", " ")
			}
		}
		fmt.Println()
	}

	time.Sleep(time.Second / 6)
}

func main() {
	board := newBoard(30)

	// simple glider
	board[1][2] = true
	board[2][3] = true
	board[3][1] = true
	board[3][2] = true
	board[3][3] = true

	// random board
	for i := range board {
		for j := range board[i] {
			board[i][j] = rand.Intn(2) == 1
		}
	}

	for i := 0; i < 60; i++ {
		render(board)
		tick(board)
	}
}
