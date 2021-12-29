package day25

import (
	"fmt"
	"strings"
)

type Tile rune

const (
	EastGoing  = '>'
	SouthGoing = 'v'
	Empty      = '.'
)

type TileMap [][]rune

func parseInput(input string) TileMap {
	var tiles TileMap
	for _, line := range strings.Split(input, "\n") {
		var tileRow []rune
		for _, r := range line {
			tileRow = append(tileRow, r)
		}
		tiles = append(tiles, tileRow)
	}
	return tiles
}

func Part1(input string) int {
	tiles := parseInput(input)
	fmt.Printf("%q\n", tiles)
	return len(tiles)
}

func Part2(input string) int {
	tiles := parseInput(input)
	return len(tiles)
}
