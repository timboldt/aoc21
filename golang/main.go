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
	"github.com/timboldt/aoc21/day06"
	"github.com/timboldt/aoc21/day07"
	"github.com/timboldt/aoc21/day08"
	"github.com/timboldt/aoc21/day09"
	"github.com/timboldt/aoc21/day10"
	"github.com/timboldt/aoc21/day11"
	"github.com/timboldt/aoc21/day12"
	"github.com/timboldt/aoc21/day13"
	"github.com/timboldt/aoc21/day14"
	"github.com/timboldt/aoc21/day15"
	"github.com/timboldt/aoc21/day16"
	"github.com/timboldt/aoc21/day17"
	"github.com/timboldt/aoc21/day18"
	"github.com/timboldt/aoc21/day19"
	"github.com/timboldt/aoc21/day20"
	"github.com/timboldt/aoc21/day21"
	"github.com/timboldt/aoc21/day22"
	"github.com/timboldt/aoc21/day23"
	"github.com/timboldt/aoc21/day24"
	"github.com/timboldt/aoc21/day25"
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
	case 6:
		part1 = day06.Part1(problemInput)
		part2 = day06.Part2(problemInput)
	case 7:
		part1 = day07.Part1(problemInput)
		part2 = day07.Part2(problemInput)
	case 8:
		part1 = day08.Part1(problemInput)
		part2 = day08.Part2(problemInput)
	case 9:
		part1 = day09.Part1(problemInput)
		part2 = day09.Part2(problemInput)
	case 10:
		part1 = day10.Part1(problemInput)
		part2 = day10.Part2(problemInput)
	case 11:
		part1 = day11.Part1(problemInput)
		part2 = day11.Part2(problemInput)
	case 12:
		part1 = day12.Part1(problemInput)
		part2 = day12.Part2(problemInput)
	case 13:
		part1 = day13.Part1(problemInput)
		part2 = day13.Part2(problemInput)
	case 14:
		part1 = day14.Part1(problemInput)
		part2 = day14.Part2(problemInput)
	case 15:
		part1 = day15.Part1(problemInput)
		part2 = day15.Part2(problemInput)
	case 16:
		part1 = day16.Part1(problemInput)
		part2 = day16.Part2(problemInput)
	case 17:
		part1 = day17.Part1(problemInput)
		part2 = day17.Part2(problemInput)
	case 18:
		part1 = day18.Part1(problemInput)
		part2 = day18.Part2(problemInput)
	case 19:
		part1 = day19.Part1(problemInput)
		part2 = day19.Part2(problemInput)
	case 20:
		part1 = day20.Part1(problemInput)
		part2 = day20.Part2(problemInput)
	case 21:
		part1 = day21.Part1(problemInput)
		part2 = day21.Part2(problemInput)
	case 22:
		part1 = day22.Part1(problemInput)
		part2 = day22.Part2(problemInput)
	case 23:
		part1 = day23.Part1(problemInput)
		part2 = day23.Part2(problemInput)
	case 24:
		part1 = day24.Part1(problemInput)
		part2 = day24.Part2(problemInput)
	case 25:
		part1 = day25.Part1(problemInput)
		part2 = day25.Part2(problemInput)
	default:
		fmt.Printf("Day %02d is not implemented yet.\n", day)
		os.Exit(1)
	}
	fmt.Printf("Solutions for Day %02d:\n"+
		"\tPart 1:\t%d\n"+
		"\tPart 2:\t%d\n", day, part1, part2)
}
