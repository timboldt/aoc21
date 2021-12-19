package day03

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, int) {
	var vals []int
	var bits int
	for _, line := range strings.Split(input, "\n") {
		if bits == 0 {
			bits = len(line)
		}
		i, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			// Should never happen.
			panic(err)
		}
		vals = append(vals, int(i))
	}
	return vals, bits
}

func Part1(input string) int {
	vals, bits := parseInput(input)
	var gamma int
	for bit := 0; bit < bits; bit++ {
		mask := 1 << bit
		cnt := 0
		for _, val := range vals {
			if val&mask != 0 {
				cnt++
			}
		}
		if cnt > len(vals)/2 {
			gamma += mask
		}
	}
	// Epsilon is just a bit flip of Gamma, of the appropriate shortened integer size.
	epsilon := gamma ^ ((1 << bits) - 1)
	return gamma * epsilon
}

func Part2(input string) int {
	vals, bits := parseInput(input)
	var oxyMask, scrubMask int
	for bit := 0; bit < bits; bit++ {
		mask := 1 << bit
		cnt := 0
		for _, val := range vals {
			if val&mask != 0 {
				cnt++
			}
		}
		if cnt >= len(vals)/2 {
			oxyMask += mask
		}
		if cnt <= len(vals)/2 {
			scrubMask += mask
		}
	}
	// fmt.Println(oxyMask)
	// fmt.Println(scrubMask)

	var oxy, scrub int
	oxyResid := math.MaxInt
	scrubResid := math.MaxInt
	for _, val := range vals {
		var resid int
		resid = val ^ oxyMask
		if resid < oxyResid {
			oxy = val
			oxyResid = resid
		}
		resid = val ^ scrubMask
		if resid < scrubResid {
			scrub = val
			scrubResid = resid
		}
	}
	// fmt.Println(oxy)
	// fmt.Println(scrub)

	return oxy * scrub
}
