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
)

func main() {
	// Parse arguments.
	if len(os.Args) != 2 {
		fmt.Println("Usage: aoc21 <day>")
		os.Exit(1)
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is an invalid day number.\n", os.Args[1])
		os.Exit(1)
	}

	// Read the input file and clean it up.
	file, err := os.Open(fmt.Sprintf("day%02d/input.txt", day))
	if err != nil {
		fmt.Printf("Cannot open input file for Day %02d.\n", day)
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
	default:
		fmt.Printf("Day %02d is not implemented yet.\n", day)
		os.Exit(1)
	}
	fmt.Printf("Solutions for Day %02d:\n"+
		"\tPart 1:\t%d\n"+
		"\tPart 2:\t%d\n", day, part1, part2)
}
