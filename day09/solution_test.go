package day09

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
2199943210
3987894921
9856789892
8767896789
9899965678`, "\n")

func TestPart1(t *testing.T) {
	want := 15
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 1134
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
