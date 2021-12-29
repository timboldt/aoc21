package day11

import (
	"strings"
)

type Position struct {
	row int
	col int
}

func flash(pos Position, octos *[10][10]int, flashes *int) {
	if octos[pos.row][pos.col] != 10 {
		return
	}
	*flashes++
	for row := pos.row - 1; row <= pos.row+1; row++ {
		for col := pos.col - 1; col <= pos.col+1; col++ {
			if row < 0 || row > 9 || col < 0 || col > 9 {
				continue
			}
			if row == pos.row && col == pos.col {
				continue
			}
			octos[row][col]++
			flash(Position{row, col}, octos, flashes)
		}
	}
}

func parseInput(input string) [10][10]int {
	var octos [10][10]int
	for row, line := range strings.Split(input, "\n") {
		for col, r := range line {
			octos[row][col] = int(r - '0')
		}
	}
	return octos
}

func Part1(input string) int {
	octos := parseInput(input)
	var flashes int
	for step := 0; step < 100; step++ {
		// Increment by 1.
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				octos[row][col]++
				if octos[row][col] == 10 {
					flash(Position{row, col}, &octos, &flashes)
				}
			}
		}
		// Reset flashed.
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				if octos[row][col] > 9 {
					octos[row][col] = 0
				}
			}
		}
	}
	return flashes
}

func Part2(input string) int {
	octos := parseInput(input)
	for step := 1; ; step++ {
		var flashes int
		// Increment by 1.
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				octos[row][col]++
				if octos[row][col] == 10 {
					flash(Position{row, col}, &octos, &flashes)
				}
			}
		}
		if flashes == 100 {
			return step
		}
		// Reset flashed.
		for row := 0; row < 10; row++ {
			for col := 0; col < 10; col++ {
				if octos[row][col] > 9 {
					octos[row][col] = 0
				}
			}
		}
	}
}
