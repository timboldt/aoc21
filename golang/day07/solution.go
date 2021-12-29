package day07

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, int, int) {
	var crabs []int
	min := math.MaxInt
	max := math.MinInt
	for _, raw := range strings.Split(input, ",") {
		crab, err := strconv.Atoi(raw)
		if err != nil {
			// Should never happen.
			panic(err)
		}
		crabs = append(crabs, crab)
		if crab < min {
			min = crab
		}
		if crab > max {
			max = crab
		}
	}
	return crabs, max, min
}

func Part1(input string) int {
	crabs, max, min := parseInput(input)
	best := math.MaxInt
	for c := min; c <= max; c++ {
		var sum int
		for _, crab := range crabs {
			d := crab - c
			if d < 0 {
				d = -d
			}
			sum += d
		}
		if sum < best {
			best = sum
		}
	}
	return best
}

func Part2(input string) int {
	crabs, max, min := parseInput(input)
	best := math.MaxInt
	for c := min; c <= max; c++ {
		var sum int
		for _, crab := range crabs {
			d := crab - c
			if d < 0 {
				d = -d
			}
			// Triangle number series.
			sum += d * (d + 1) / 2
		}
		if sum < best {
			best = sum
		}
	}
	return best
}
