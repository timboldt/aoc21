package day24

import (
	"strings"
)

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func Part1(input string) int {
	lines := parseInput(input)
	return len(lines)
}

func Part2(input string) int {
	lines := parseInput(input)
	return len(lines)
}
