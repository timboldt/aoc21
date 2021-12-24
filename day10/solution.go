package day10

import (
	"sort"
	"strings"
)

type ParseResult int

const (
	ParseOk ParseResult = iota
	ParseTruncated
	ParseEndChunk
	ParseFailed
)

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func chunkEnd(r byte) byte {
	switch r {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	default:
		return 0
	}
}

func scorePart1(r byte) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}

func scorePart2(r byte) int {
	switch r {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	default:
		return 0
	}
}

func validateChunk(s []byte) (int, ParseResult, []byte) {
	startIdx := 0
	endIdx := 1
	for {
		if startIdx == len(s) {
			// Success.
			return len(s), ParseOk, nil
		}
		endWanted := chunkEnd(s[startIdx])
		if endWanted == 0 {
			// No more valid chunks.
			return startIdx, ParseEndChunk, nil
		}
		if endIdx >= len(s) {
			// Truncated before finding end character.
			return len(s), ParseTruncated, []byte{endWanted}
		}
		if s[endIdx] == endWanted {
			// Start a new top-level chunk
			startIdx = endIdx + 1
			endIdx = startIdx + 1
			continue
		}
		subLen, result, missing := validateChunk(s[endIdx:])
		if result == ParseTruncated {
			return endIdx + subLen, ParseTruncated, append(missing, endWanted)
		}
		if subLen == 0 || result == ParseFailed {
			return endIdx + subLen, ParseFailed, nil
		}
		endIdx += subLen
	}
}

func Part1(input string) int {
	lines := parseInput(input)
	var sum int
	for _, line := range lines {
		b := []byte(line)
		idx, _, _ := validateChunk(b)
		if idx < len(b) {
			sum += scorePart1(b[idx])
		}
	}
	return sum
}

func Part2(input string) int {
	lines := parseInput(input)
	var sums []int
	for _, line := range lines {
		b := []byte(line)
		_, _, missing := validateChunk(b)
		var lineSum int
		for _, r := range missing {
			lineSum *= 5
			lineSum += scorePart2(r)
		}
		if lineSum > 0 {
			sums = append(sums, lineSum)
		}
	}
	sort.Ints(sums)
	return sums[len(sums)/2]
}
