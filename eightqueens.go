package piscine

import (
	"github.com/01-edu/z01"
)

const size = 8

var board [size][size]bool

func goodDirection(x, y, vx, vy int) bool {

	for 0 <= x && x < size &&
		0 <= y && y < size {
		if board[x][y] {

			return false
		}
		x = x + vx
		y = y + vy
	}

	return true
}

func goodSquare(x, y int) bool {
	return goodDirection(x, y, +0, -1) &&
		goodDirection(x, y, +1, -1) &&
		goodDirection(x, y, +1, +0) &&
		goodDirection(x, y, +1, +1) &&
		goodDirection(x, y, +0, +1) &&
		goodDirection(x, y, -1, +1) &&
		goodDirection(x, y, -1, +0) &&
		goodDirection(x, y, -1, -1)
}

func printQueens() {
	x := 0
	for x < size {
		y := 0
		for y < size {
			if board[x][y] {

				z01.PrintRune(rune(y) + '1')
			}
			y++
		}
		x++
	}
	z01.PrintRune('\n')
}

func tryX(x int) {
	y := 0
	for y < size {
		if goodSquare(x, y) {

			board[x][y] = true

			if x == size-1 {

				printQueens()
			} else {

				tryX(x + 1)
			}

			board[x][y] = false
		}
		y++
	}
}

func EightQueens() {

	tryX(0)
}
