package day14

import (
	"fmt"
	"math"
	"strings"
)

type Rule struct {
	pair PolyPair
	ins  byte
}

type PolyPair struct {
	first, second byte
}

type Polymer struct {
	pairs map[PolyPair]int
	last  byte
}

func parseInput(input string) (Polymer, []Rule) {
	var (
		poly  Polymer
		rules []Rule
	)
	poly.pairs = make(map[PolyPair]int)
	for idx, line := range strings.Split(input, "\n") {
		if idx == 0 {
			var prev byte
			for _, c := range line {
				if prev != 0 {
					poly.pairs[PolyPair{first: prev, second: byte(c)}]++
				}
				prev = byte(c)
			}
			poly.last = prev
			continue
		}
		if idx == 1 {
			continue
		}
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			panic(fmt.Errorf("invalid rule: %q", line))
		}
		rules = append(rules, Rule{
			pair: PolyPair{
				first:  parts[0][0],
				second: parts[0][1],
			},
			ins: parts[1][0],
		})
	}
	return poly, rules
}

func applyRules(p *Polymer, rules []Rule) {
	out := make(map[PolyPair]int)
OUTER:
	for k, v := range p.pairs {
		for _, r := range rules {
			if k == r.pair {
				out[PolyPair{k.first, r.ins}] += v
				out[PolyPair{r.ins, k.second}] += v
				continue OUTER
			}
		}
		out[k] += v
	}
	p.pairs = out
}

func score(p Polymer) int {
	// Count up the starting values of all of the pairs.
	cnt := make(map[byte]int)
	for k, v := range p.pairs {
		cnt[k.first] += v
	}
	// Add in the last element too.
	cnt[p.last]++

	min := math.MaxInt
	max := math.MinInt
	for _, v := range cnt {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return max - min
}

func Part1(input string) int {
	poly, rules := parseInput(input)
	fmt.Println(poly)
	for i := 0; i < 10; i++ {
		applyRules(&poly, rules)
	}
	return score(poly)
}

func Part2(input string) int {
	poly, rules := parseInput(input)
	for i := 0; i < 40; i++ {
		applyRules(&poly, rules)
	}
	return score(poly)
}
