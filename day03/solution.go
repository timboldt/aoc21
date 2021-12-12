package day03

import (
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, int) {
	var vals []int
	var bits int
	for _, line := range strings.Split(input, "\n") {
		if bits == 0 {
			bits = len(line)
		}
		i, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			// Should never happen.
			panic(err)
		}
		vals = append(vals, int(i))
	}
	return vals, bits
}

func Part1(input string) int {
	vals, bits := parseInput(input)
	var gamma int
	for bit := 0; bit < bits; bit++ {
		mask := 1 << bit
		cnt := 0
		for _, val := range vals {
			if val&mask != 0 {
				cnt++
			}
		}
		if cnt > len(vals)/2 {
			gamma += mask
		}
	}
	// Epsilon is just a bit flip of Gamma, of the appropriate shortened integer size.
	epsilon := gamma ^ ((1 << bits) - 1)
	return gamma * epsilon
}

func Part2(input string) int {
	_, bits := parseInput(input)
	return bits
}
