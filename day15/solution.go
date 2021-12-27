package day15

import (
	"fmt"
	"strings"
)

type Point struct {
	x, y int
}

type Node struct {
	risk int
	min  int
}

func parseInput(input string) [][]Node {
	var nodes [][]Node
	for row, line := range strings.Split(input, "\n") {
		nodes = append(nodes, []Node{})
		for _, r := range line {
			nodes[row] = append(nodes[row], Node{
				risk: int(r - '0'),
				min:  99999,
			})
		}
	}
	return nodes
}

func findPath(nodes [][]Node) []Point {
	var (
	//path  []Point
	)
	nodes[0][0].min = 0
	queue := []Point{{x: 0, y: 0}}
	for {
		fmt.Println(queue)
		lastRow := queue[0].x
		lastCol := queue[0].y
		queue = queue[1:]

		for row := lastRow - 1; row <= lastRow+1; row++ {
			for col := lastCol - 1; col <= lastCol+1; col++ {
				if row < 0 || row >= len(nodes) {
					continue
				}
				if col < 0 || col >= len(nodes[0]) {
					continue
				}
				if row != 0 && col != 0 {
					continue
				}
				if row == 0 && col == 0 {
					continue
				}
				tot := nodes[lastRow][lastCol].min + nodes[row][col].risk
				fmt.Println(tot, nodes[row][col].min)
				if nodes[row][col].min > tot {
					nodes[row][col].min = tot
					queue = append(queue, Point{x: col, y: row})
				}
				if row == 9 && col == 9 {
					return nil
				}
			}
		}
		if len(queue) == 0 {
			break
		}
		// if result := findPath(nodes, append(path, Point{x: col, y: row})); result != nil {
		// 	return result
		// }
	}
	return nil
}

func Part1(input string) int {
	nodes := parseInput(input)
	path := findPath(nodes)
	return len(path)
}

func Part2(input string) int {
	lines := parseInput(input)
	return len(lines)
}
