package day07

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
16,1,2,0,4,2,7,1,2,14`, "\n")

func TestPart1(t *testing.T) {
	want := 37
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 168
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
