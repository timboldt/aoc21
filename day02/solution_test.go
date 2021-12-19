package day02

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
forward 5
down 5
forward 8
up 3
down 8
forward 2`, "\n")

func TestPart1(t *testing.T) {
	want := 150
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 900
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
