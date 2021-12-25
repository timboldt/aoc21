package day12

import (
	"strings"
	"testing"
)

var sampleTextA = strings.Trim(`
start-A
start-b
A-c
A-b
b-d
A-end
b-end`, "\n")

var sampleTextB = strings.Trim(`
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`, "\n")

var sampleTextC = strings.Trim(`
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`, "\n")

func TestPart1A(t *testing.T) {
	want := 10
	got := Part1(sampleTextA)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart1B(t *testing.T) {
	want := 19
	got := Part1(sampleTextB)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart1C(t *testing.T) {
	want := 226
	got := Part1(sampleTextC)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2A(t *testing.T) {
	want := 36
	got := Part2(sampleTextA)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2B(t *testing.T) {
	want := 103
	got := Part2(sampleTextB)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2C(t *testing.T) {
	want := 3509
	got := Part2(sampleTextC)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
