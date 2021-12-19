package day03

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`, "\n")

func TestPart1(t *testing.T) {
	want := 198
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 230
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
