package day13

import (
	"strings"
	"testing"
)

var sampleText = strings.Trim(`
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`, "\n")

func TestPart1(t *testing.T) {
	want := 17
	got := Part1(sampleText)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

// func TestPart2(t *testing.T) {
// 	want := 999
// 	got := Part2(sampleText)
// 	if got != want {
// 		t.Fatalf(`Got %d, want %d`, got, want)
// 	}
// }
