package day08

import (
	"strings"
	"testing"
)

var sampleText1 = strings.Trim(`
acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`, "\n")

var sampleText2 = strings.Trim(`
be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`, "\n")

func TestPart1(t *testing.T) {
	want := 26
	got := Part1(sampleText2)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart21(t *testing.T) {
	want := 5353
	got := Part2(sampleText1)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}

func TestPart2b(t *testing.T) {
	want := 61229
	got := Part2(sampleText2)
	if got != want {
		t.Fatalf(`Got %d, want %d`, got, want)
	}
}
