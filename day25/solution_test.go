package day25

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`, "\n")

func TestPart1(t *testing.T) {
	want := 58
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 999
	got := Part2(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
