package day06

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
3,4,3,1,2`, "\n")

func TestPart1(t *testing.T) {
	want := 5934
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 26984457539
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
