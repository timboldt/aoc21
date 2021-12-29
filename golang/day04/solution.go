package day04

import (
	"math"
	"strconv"
	"strings"
)

type Board [5][5]int

func parseInput(input string) (bingoCalls []int, boards []Board) {
	var row int
	for _, line := range strings.Split(input, "\n") {
		// Comma-filled line: record the bingo calls.
		if strings.Contains(line, ",") {
			for _, raw := range strings.Split(line, ",") {
				num, err := strconv.Atoi(raw)
				if err != nil {
					// Should never happen.
					panic(err)
				}
				bingoCalls = append(bingoCalls, num)
			}
			continue
		}
		// Blank line: start a new board.
		if len(line) == 0 {
			boards = append(boards, Board{})
			row = 0
			continue
		}
		// Parse the next row of the Bingo board (five space-separated values)
		col := 0
		for _, raw := range strings.Split(line, " ") {
			// Due to column-aligned input, there may be extraneous blank values.
			if len(raw) == 0 {
				continue
			}
			num, err := strconv.Atoi(raw)
			if err != nil {
				// Should never happen.
				panic(err)
			}
			boards[len(boards)-1][row][col] = num
			col++
		}
		row++
	}
	return
}

func computeWin(bingoCalls []int, board Board) (int, int) {
	var bingo [5][5]bool
	for idx, call := range bingoCalls {
		// Mark the called number.
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if board[row][col] == call {
					bingo[row][col] = true
				}
			}
		}
		var winner bool
		for i := 0; i < 5; i++ {
			var rowCnt, colCnt int
			for j := 0; j < 5; j++ {
				if bingo[i][j] {
					rowCnt++
				}
				if bingo[j][i] {
					colCnt++
				}
			}
			if rowCnt == 5 || colCnt == 5 {
				// Bingo!
				winner = true
				break
			}
		}
		if winner {
			// Sum the unmarked values.
			var sum int
			for row := 0; row < 5; row++ {
				for col := 0; col < 5; col++ {
					if !bingo[row][col] {
						sum += board[row][col]
					}
				}
			}
			return idx, call * sum
		}
	}
	return len(bingoCalls), 0
}

func Part1(input string) int {
	bingoCalls, boards := parseInput(input)
	bestCalls := math.MaxInt
	bestScore := 0
	for _, board := range boards {
		calls, score := computeWin(bingoCalls, board)
		if calls < bestCalls {
			bestCalls = calls
			bestScore = score
		}
	}
	return bestScore
}

func Part2(input string) int {
	bingoCalls, boards := parseInput(input)
	worstCalls := math.MinInt
	worstScore := 0
	for _, board := range boards {
		calls, score := computeWin(bingoCalls, board)
		if calls > worstCalls {
			worstCalls = calls
			worstScore = score
		}
	}
	return worstScore
}
