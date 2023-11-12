package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var deltas = [3]int{-1, 0, 1}

type board [][]bool

func newBoard(width, height int) board {
	cells := make(board, height)
	for i := range cells {
		cells[i] = make([]bool, width)
	}
	return cells
}

func copyBoard(b board) board {
	out := newBoard(len(b), len(b[0]))

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
	count := 0
	for _, i := range deltas {
		for _, j := range deltas {
			if i == 0 && j == 0 {
				continue
			}

			nY := (y + j + 10) % 10
			nX := (x + i + 10) % 10

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
	board := newBoard(10, 10)

	board[1][2] = true
	board[2][3] = true
	board[3][1] = true
	board[3][2] = true
	board[3][3] = true

	for i := 0; i < 60; i++ {
		render(board)
		tick(board)
	}
}
