package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/timboldt/aoc21/day01"
	"github.com/timboldt/aoc21/day02"
	"github.com/timboldt/aoc21/day03"
	"github.com/timboldt/aoc21/day04"
	"github.com/timboldt/aoc21/day05"
)

func main() {
	// Parse arguments.
	var fnameText, dayText string
	switch len(os.Args) {
	case 2:
		fnameText = "input.txt"
		dayText = os.Args[1]
	case 3:
		fnameText = os.Args[1]
		dayText = os.Args[2]
	default:
		fmt.Println("Usage: aoc21 [filename] day")
		os.Exit(1)
	}
	day, err := strconv.Atoi(dayText)
	if err != nil {
		fmt.Printf("%s is an invalid day number.\n", dayText)
		os.Exit(1)
	}

	// Read the input file and clean it up.
	fname := fmt.Sprintf("day%02d/%s", day, fnameText)
	file, err := os.Open(fname)
	if err != nil {
		fmt.Printf("Cannot open input file '%s' for Day %02d.\n", fname, day)
		os.Exit(1)
	}
	defer file.Close()
	rawInput, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		fmt.Printf("Cannot read input file for Day %02d.\n", day)
		os.Exit(1)
	}
	problemInput := strings.TrimSuffix(string(rawInput), "\n")

	var part1, part2 int
	switch day {
	case 1:
		part1 = day01.Part1(problemInput)
		part2 = day01.Part2(problemInput)
	case 2:
		part1 = day02.Part1(problemInput)
		part2 = day02.Part2(problemInput)
	case 3:
		part1 = day03.Part1(problemInput)
		part2 = day03.Part2(problemInput)
	case 4:
		part1 = day04.Part1(problemInput)
		part2 = day04.Part2(problemInput)
	case 5:
		part1 = day05.Part1(problemInput)
		part2 = day05.Part2(problemInput)
	default:
		fmt.Printf("Day %02d is not implemented yet.\n", day)
		os.Exit(1)
	}
	fmt.Printf("Solutions for Day %02d:\n"+
		"\tPart 1:\t%d\n"+
		"\tPart 2:\t%d\n", day, part1, part2)
}
