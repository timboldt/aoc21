package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: aoc21 <day>")
		os.Exit(1)
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%s is an invalid day number\n", os.Args[1])
		os.Exit(1)
	}
	file, err := os.Open(fmt.Sprintf("day%02d/input.txt", day))
	if err != nil {
		fmt.Printf("Cannot open input file for Day %02d\n", day)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	rawInput, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Printf("Cannot read input file for Day %02d\n", day)
		os.Exit(1)
	}
	problemInput := strings.TrimSuffix(string(rawInput), "\n")
	fmt.Println(problemInput)
}
