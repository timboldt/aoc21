package day01

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
199
200
208
210
200
207
240
269
260
263`, "\n")

func TestPart1(t *testing.T) {
	want := 7
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 5
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
