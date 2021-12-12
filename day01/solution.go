package day01

import (
	"math"
	"strconv"
	"strings"
)

func Part1(input string) int {
	var vals []int
	for _, line := range strings.Split(input, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			// Should never happen.
			panic(err)
		}
		vals = append(vals, i)
	}
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
	return 42
}
