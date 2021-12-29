package day01

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	var vals []int
	for _, line := range strings.Split(input, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			// Should never happen.
			panic(err)
		}
		vals = append(vals, i)
	}
	return vals
}

func Part1(input string) int {
	vals := parseInput(input)
	prevVal := math.MaxInt
	increases := 0
	for _, val := range vals {
		if val > prevVal {
			increases++
		}
		prevVal = val
	}
	return increases
}

func Part2(input string) int {
	vals := parseInput(input)
	prevVal := math.MaxInt
	increases := 0
	for idx := 0; idx < len(vals)-2; idx++ {
		val := vals[idx] + vals[idx+1] + vals[idx+2]
		if val > prevVal {
			increases++
		}
		prevVal = val
	}
	return increases
}
